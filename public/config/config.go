package config

var cfg = &Cfg{}

type Cfg struct {
	Http  HttpConfig  `env:"HTTP"`
	Redis RedisConfig `env:"REDIS"`
}

type HttpConfig struct {
	Port         int `env:"PORT,default=8080"`
	GraceTimeout int `env:"GRACE_TIMEOUT,default=5"`
	ReadTimeout  int `env:"READ_TIMEOUT,default=0"`
	WriteTimeout int `env:"WRITE_TIMEOUT,default=0"`
}

type RedisConfig struct {
	Address     string `env:"ADDRESS"`
	Password    string `env:"PASSWORD"`
	DB          int    `env:"DB,default=0"`
	PoolSize    int    `env:"POOL_SIZE,default=3"`
	MaxRetries  int    `env:"MAX_RETRIES,default=1"`
	MinIdleSize int    `env:"MIN_IDLE_SIZE,default=1"`
}

func Config() *Cfg {
	return cfg
}

func init() {
	if err := ParseEnv(cfg); err != nil {
		panic(err)
	}
}
