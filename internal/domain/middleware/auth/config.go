package auth

func NewDefaultConfig() *Config {
	return &Config{
		secret:           "dogo",
		Expire:           3600,
		AuthKey:          "Authorization",
		SessionKey:       "jwt_token",
		SessionKeyExpire: 3600,
	}
}

type ConfigFunc func(cfg *Config)

type Config struct {
	secret            string
	Expire            int64
	ignoreRoute       []string
	ignorePrefix      []string
	ignoreSuffix      []string
	ignoreMethodRoute [][2]string
	AuthKey           string
	SessionKey        string
	SessionKeyExpire  int
}

func Secret(secret string) ConfigFunc {
	return func(cfg *Config) {
		cfg.secret = secret
	}
}

func Expire(expire int64) ConfigFunc {
	return func(cfg *Config) {
		cfg.Expire = expire
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
		cfg.AuthKey = authKey
	}
}

func SessionKey(sessionKey string) ConfigFunc {
	return func(cfg *Config) {
		cfg.SessionKey = sessionKey
	}
}

func SessionKeyExpire(sessionKeyExpire int) ConfigFunc {
	return func(cfg *Config) {
		cfg.SessionKeyExpire = sessionKeyExpire
	}
}
