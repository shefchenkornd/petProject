package store

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"petProject/config"
	"petProject/logger"
	"petProject/store/mysql"
	"time"
)

// KeepAlivePollPeriod is a Pg/MySQL keepalive check time period
const KeepAlivePollPeriod = 3

type Store struct {
	MySQL *mysql.MySQL // for KeepAliveMySQL (see below)

	User UserRepo
}

// New creates new store
func New(ctx context.Context) (*Store, error) {
	cfg := config.Get()

	// connect to MySQL
	mysqlDB, err := mysql.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "mysqldb.Dial failed")
	}

	// Run MySQL migrations
	if mysqlDB != nil {
		log.Println("Running MySQL migrations...")
		if err := runMysqlMigrations(); err != nil {
			return nil, errors.Wrap(err, "runMysqlMigrations failed")
		}
	}

	var store Store

	// Init MySQL repositories
	if mysqlDB != nil {
		store.MySQL = mysqlDB
		go store.KeepAliveMySQL()
		store.User = mysql.NewUserRepo(mysqlDB)
	}

	return &store, nil
}

// KeepAliveMySQL makes sure MySQL is alive and reconnects if needed
func (store *Store) KeepAliveMySQL() {
	logger := logger.Get()
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if store.MySQL == nil {
			lostConnect = true
		} else if err = store.MySQL.DB.DB().Ping(); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		logger.Debug().Msg("[store.KeepAliveMySQL] Lost MySQL connection. Restoring...")
		store.MySQL, err = mysql.Dial()
		if err != nil {
			logger.Err(err)
			continue
		}
		logger.Debug().Msg("[store.KeepAliveMySQL] MySQL reconnected")
	}
}
