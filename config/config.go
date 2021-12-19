package config

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

type Config struct {
	LogLevel            string `envconfig:"LOG_LEVEL"`
	HTTPAddr            string `envconfig:"HTTP_ADDR"`
	PgAddr              string `envconfig:"PG_ADDR"`
	PgMigrationsPath    string `envconfig:"PG_MIGRATIONS_PATH"`
	MysqlAddr           string `envconfig:"MYSQL_ADDR"`
	MysqlUser           string `envconfig:"MYSQL_USER"`
	MysqlPassword       string `envconfig:"MYSQL_PASSWORD"`
	MysqlDB             string `envconfig:"MYSQL_DB"`
	MysqlMigrationsPath string `envconfig:"MYSQL_MIGRATIONS_PATH"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		configBytes, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
