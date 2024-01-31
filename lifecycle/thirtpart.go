package lifecycle

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	MySQL *gorm.DB
	Redis *redis.Client
)

func InjectMySQL(db interface{}) {
	switch db.(type) {
	case *gorm.DB:
		MySQL = db.(*gorm.DB)
	case func() *gorm.DB:
		MySQL = db.(func() *gorm.DB)()
	default:
		panic("invalid mysql type")
	}
}

func InjectRedis(rds interface{}) {
	switch rds.(type) {
	case *redis.Client:
		Redis = rds.(*redis.Client)
	case func() *redis.Client:
		Redis = rds.(func() *redis.Client)()
	default:
		panic("invalid redis type")
	}
}
