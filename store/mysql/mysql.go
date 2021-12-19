package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"petProject/config"
)

type Mysql struct {
	*gorm.DB
}

func Dial() (*Mysql, error) {
	cfg := config.Get()
	if cfg.MysqlDB == "" {
		return nil, nil
	}

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlAddr, cfg.MysqlDB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Mysql{db}, nil
}
