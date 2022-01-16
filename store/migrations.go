package store

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"petProject/config"
)

// runMysqlMigrations runs MySQL migrations
func runMysqlMigrations() error {
	cfg := config.Get()
	if cfg.MysqlMigrationsPath == "" {
		return nil
	}
	if cfg.MysqlDB == "" {
		return errors.New("No cfg.MysqlDB provided")
	}

	m, err := migrate.New(
		cfg.MysqlMigrationsPath,
		fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s", cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlAddr, cfg.MysqlDB),
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
