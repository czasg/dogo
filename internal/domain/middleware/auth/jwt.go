package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"proj/public/utils"
	"strings"
	"time"
)

var (
	JwtKey                  = "jwt"
	JwtPayloadKey           = "jwt:payload"
	LogoutKeyTmpl           = "jwt:logout:uid:%d"
	JwtTokenInvalidErr      = errors.New("invalid jwt token")
	JwtTokenVerificationErr = errors.New("jwt verification failure")
	JwtTokenExpireErr       = errors.New("token expire")
)

func NewJwtHandler(cache *redis.Client, configFunc ...ConfigFunc) gin.HandlerFunc {
	config := NewDefaultConfig()
	for _, patch := range configFunc {
		patch(config)
	}
	jwt := &Jwt{
		Config: config,
	}
	return func(c *gin.Context) {
		c.Set(JwtKey, jwt)
		method := c.Request.Method
		route := c.Request.URL.Path
		for _, ignoreRoute := range jwt.Config.ignoreRoute {
			if route == ignoreRoute {
				c.Next()
				return
			}
		}
		for _, ignorePrefix := range jwt.Config.ignorePrefix {
			if strings.HasPrefix(route, ignorePrefix) {
				c.Next()
				return
			}
		}
		for _, ignoreSuffix := range jwt.Config.ignoreSuffix {
			if strings.HasSuffix(route, ignoreSuffix) {
				c.Next()
				return
			}
		}
		for _, ignoreMethodRoute := range jwt.Config.ignoreMethodRoute {
			if method == ignoreMethodRoute[0] && route == ignoreMethodRoute[1] {
				c.Next()
				return
			}
		}
		token := c.GetHeader(jwt.Config.AuthKey)
		if token == "" {
			token, _ = c.Cookie(jwt.Config.SessionKey)
		}
		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimSpace(token[7:])
		}
		payload, err := jwt.Valid(token)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		lastLogoutTime, err := cache.Get(c, fmt.Sprintf(LogoutKeyTmpl, payload.UserID)).Int64()
		if err != nil && !errors.Is(err, redis.Nil) {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if lastLogoutTime > 0 && lastLogoutTime > payload.Issue {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Set(JwtPayloadKey, payload)
		c.Next()
	}
}

type Jwt struct {
	H      utils.Hash
	Config *Config
}

type JwtPayload struct {
	UserName string `json:"une"`
	UserID   int64  `json:"uid"`
	Expire   int64  `json:"exp"`
	Issue    int64  `json:"iss"`
	Admin    bool   `json:"admin"`
}

func (j *Jwt) EncryptPayload(payload JwtPayload) (string, error) {
	now := time.Now().Unix()
	payload.Issue = now
	payload.Expire = now + j.Config.Expire
	payloadB64, err := j.encrypt(payload)
	if err != nil {
		return "", err
	}
	// header
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	headerB64, err := j.encrypt(header)
	if err != nil {
		return "", err
	}
	// secret
	sb := strings.Builder{}
	sb.Write([]byte(headerB64))
	sb.Write([]byte("."))
	sb.Write([]byte(payloadB64))
	sb.Write([]byte("."))
	sb.WriteString(j.H.Sha256([]byte(headerB64), []byte(payloadB64)))
	return sb.String(), nil
}

func (j *Jwt) Encrypt(name string, id int64, isAdmin bool) (string, error) {
	return j.EncryptPayload(JwtPayload{
		UserName: name,
		UserID:   id,
		Admin:    isAdmin,
	})
}

func (j *Jwt) encrypt(data interface{}) (string, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(body), nil
}

func (j Jwt) Decrypt(token string) (*JwtPayload, error) {
	ss := strings.Split(token, ".")
	if len(ss) != 3 {
		return nil, JwtTokenInvalidErr
	}
	payloadBody, err := base64.StdEncoding.DecodeString(ss[1])
	if err != nil {
		return nil, err
	}
	payload := JwtPayload{}
	err = json.Unmarshal(payloadBody, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (j *Jwt) Valid(token string) (*JwtPayload, error) {
	ss := strings.Split(token, ".")
	if len(ss) != 3 {
		return nil, JwtTokenInvalidErr
	}
	if j.H.Sha256([]byte(ss[0]), []byte(ss[1])) != ss[2] {
		return nil, JwtTokenVerificationErr
	}
	payload, err := j.Decrypt(token)
	if err != nil {
		return nil, JwtTokenInvalidErr
	}
	if payload.Issue > time.Now().Unix() {
		return nil, JwtTokenExpireErr
	}
	if payload.Expire < time.Now().Unix() {
		return nil, JwtTokenExpireErr
	}
	return payload, nil
}

type JwtService struct {
	Cache *redis.Client
}

func (j *JwtService) Enable(c *gin.Context) (*Jwt, *JwtPayload, bool) {
	jwtVal, ok := c.Get(JwtKey)
	if !ok {
		return nil, nil, false
	}
	jwt, ok := jwtVal.(*Jwt)
	if !ok {
		return nil, nil, false
	}
	payloadVal, ok := c.Get(JwtPayloadKey)
	if !ok {
		return nil, nil, false
	}
	payload, ok := payloadVal.(*JwtPayload)
	if !ok {
		return nil, nil, false
	}
	return jwt, payload, true
}

func (j *JwtService) Logout(c *gin.Context, jwt *Jwt, payload *JwtPayload) error {
	return j.Cache.Set(c, fmt.Sprintf(LogoutKeyTmpl, payload.UserID), time.Now().Unix(), time.Duration(jwt.Config.Expire)*time.Second).Err()
}
