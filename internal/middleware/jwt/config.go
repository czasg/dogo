package jwt

func NewDefaultConfig() *Config {
	return &Config{
		secret:     "dogo",
		expire:     3600,
		authKey:    "Authorization",
		sessionKey: "jwt_token",
	}
}

type ConfigFunc func(cfg *Config)

type Config struct {
	secret            string
	expire            int64
	ignoreRoute       []string
	ignorePrefix      []string
	ignoreSuffix      []string
	ignoreMethodRoute [][2]string
	authKey           string
	sessionKey        string
}

func Secret(secret string) ConfigFunc {
	return func(cfg *Config) {
		cfg.secret = secret
	}
}

func Expire(expire int64) ConfigFunc {
	return func(cfg *Config) {
		cfg.expire = expire
	}
}

func IgnoreRoute(route []string) ConfigFunc {
	return func(cfg *Config) {
		cfg.ignoreRoute = route
	}
}

func IgnorePrefix(prefix []string) ConfigFunc {
	return func(cfg *Config) {
		cfg.ignorePrefix = prefix
	}
}

func IgnoreSuffix(suffix []string) ConfigFunc {
	return func(cfg *Config) {
		cfg.ignoreSuffix = suffix
	}
}

func IgnoreMethodRoute(mr [][2]string) ConfigFunc {
	return func(cfg *Config) {
		cfg.ignoreMethodRoute = mr
	}
}

func AuthKey(authKey string) ConfigFunc {
	return func(cfg *Config) {
		cfg.authKey = authKey
	}
}

func SessionKey(sessionKey string) ConfigFunc {
	return func(cfg *Config) {
		cfg.sessionKey = sessionKey
	}
}
