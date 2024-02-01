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
		case *gorm.DB:
			MySQL = ins
		case func() *gorm.DB:
			MySQL = ins()
		case *redis.Client:
			Redis = ins
		case func() *redis.Client:
			Redis = ins()
		default:
			panic("invalid inject type")
		}
	}
}
