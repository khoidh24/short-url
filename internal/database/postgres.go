package database

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/khoidh24/short-url/internal/config"
)

var (
	db   *gorm.DB
	once sync.Once
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		dsn := config.LoadConfig().DBUrl
		var err error
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database: ", err)
		}
	})
	return db
}
