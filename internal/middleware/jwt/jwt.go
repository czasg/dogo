package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"proj/public/utils"
	"strings"
	"time"
)

var (
	JwtTokenInvalidErr      = errors.New("invalid jwt token")
	JwtTokenVerificationErr = errors.New("jwt verification failure")
)

func NewJwtHandler(configFunc ...ConfigFunc) gin.HandlerFunc {
	config := NewDefaultConfig()
	for _, patch := range configFunc {
		patch(config)
	}
	jwt := Jwt{
		Config: config,
	}
	return func(c *gin.Context) {
		c.Set("jwt", jwt)
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
		token := c.GetHeader(jwt.Config.authKey)
		if token == "" {
			token, _ = c.Cookie(jwt.Config.sessionKey)
		}
		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		err := jwt.Valid(token)
		if err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}

type Jwt struct {
	H      utils.Hash
	Config *Config
}

type JWTPayload map[string]interface{}

func (j Jwt) Encrypt(data JWTPayload) ([]byte, error) {
	// payload
	if data == nil {
		data = make(JWTPayload)
	}
	data["exp"] = time.Now().Unix() + j.Config.expire
	payloadB64, err := j.encrypt(data)
	if err != nil {
		return nil, err
	}
	// header
	header := JWTPayload{
		"alg": "HS256",
		"typ": "JWT",
	}
	headerB64, err := j.encrypt(header)
	if err != nil {
		return nil, err
	}
	// secret
	sb := strings.Builder{}
	sb.Write([]byte(headerB64))
	sb.Write([]byte("."))
	sb.Write([]byte(payloadB64))
	sb.Write([]byte("."))
	sb.WriteString(j.H.Sha256([]byte(headerB64), []byte(payloadB64)))
	return []byte(sb.String()), nil
}

func (j Jwt) encrypt(data map[string]interface{}) (string, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(body), nil
}

func (j Jwt) Decrypt(token string) (JWTPayload, error) {
	ss := strings.Split(token, ".")
	if len(ss) != 3 {
		return nil, JwtTokenInvalidErr
	}
	payloadBody, err := base64.StdEncoding.DecodeString(ss[1])
	if err != nil {
		return nil, err
	}
	payload := make(JWTPayload)
	err = json.Unmarshal(payloadBody, &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (j Jwt) Valid(token string) error {
	ss := strings.Split(token, ".")
	if len(ss) != 3 {
		return JwtTokenInvalidErr
	}
	if j.H.Sha256([]byte(ss[0]), []byte(ss[1])) != ss[2] {
		return JwtTokenVerificationErr
	}
	return nil
}
