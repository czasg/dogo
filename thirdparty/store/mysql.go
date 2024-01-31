package store

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"proj/public/config"
	"time"
)

func NewMySQL(ctx context.Context, cfg config.MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Address,
		cfg.DB,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.PoolMaxIdle)
	sqlDB.SetMaxOpenConns(cfg.PoolMaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.PoolMaxLifeTime) * time.Second)
	err = sqlDB.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
