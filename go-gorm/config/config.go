package config

import (
	"gorm.io/gorm"
	"time"
)

func New(dialect gorm.Dialector) (*gorm.DB, error) {
	orm, err := gorm.Open(dialect, &gorm.Config{})
	if nil != err {
		return nil, err
	}

	sqlDB, err := orm.DB()
	if nil != err {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(300 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)

	return orm, nil
}
