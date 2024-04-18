package lifecycle

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MySQLCaller func() *gorm.DB
type RedisCaller func() *redis.Client

var (
	MySQL *gorm.DB
	Redis *redis.Client
)

func Inject(objs ...interface{}) {
	for _, obj := range objs {
		switch caller := obj.(type) {
		case MySQLCaller:
			MySQL = caller()
		case RedisCaller:
			Redis = caller()
		default:
			panic("invalid inject type")
		}
	}
}
