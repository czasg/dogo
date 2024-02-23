package lifecycle

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	MySQL *gorm.DB
	Redis *redis.Client
)

func Inject(objs ...interface{}) {
	for _, obj := range objs {
		switch ins := obj.(type) {
		case func() *gorm.DB:
			SetMySQL(ins())
		case func() *redis.Client:
			SetRedis(ins())
		default:
			panic("invalid inject type")
		}
	}
}

func SetMySQL(ins *gorm.DB) {
	MySQL = ins
}

func SetRedis(ins *redis.Client) {
	Redis = ins
}
