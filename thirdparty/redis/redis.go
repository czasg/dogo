package redis

type Config struct {
	Address     string `env:"ADDRESS"`
	Password    string `env:"PASSWORD"`
	DB          int    `env:"DB,default=0"`
	PoolSize    int    `env:"POOL_SIZE,default=3"`
	MaxRetries  int    `env:"MAX_RETRIES,default=1"`
	MinIdleSize int    `env:"MIN_IDLE_SIZE,default=1"`
}

//func NewRedis(cfg Config) (*redis.Client, error) {
//	ins := redis.NewClient(&redis.Options{
//		Addr:         cfg.Address,
//		Password:     cfg.Password,
//		DB:           cfg.DB,
//		PoolSize:     cfg.PoolSize,
//		MaxRetries:   cfg.MaxRetries,
//		MinIdleConns: cfg.MinIdleSize,
//	})
//	err := ins.Ping().Err()
//	if err != nil {
//		return nil, err
//	}
//	return ins, nil
//}
