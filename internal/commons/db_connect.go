package commons

import (
	"log"

	"github.com/hiimlyn/go-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	if cfg.DBDns == "" {
		log.Fatal("DB_URL is not set")
	}
	//connect database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DBDns,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}
