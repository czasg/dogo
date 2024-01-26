package config

var cfg *Config

type Config struct {
	Http HttpConfig
}

type HttpConfig struct {
	Port         int `env:",default=8080"`
	GraceTimeout int `env:",default=5"`
	ReadTimeout  int `env:",default=0"`
	WriteTimeout int `env:",default=0"`
}

func GetConfig() *Config {
	return cfg
}

func init() {
	config := Config{}
	if err := ParseEnv(&config); err != nil {
		panic(err)
	}
	cfg = &config
}
