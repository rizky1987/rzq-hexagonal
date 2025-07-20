package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresDBFromEnv loads .env and returns a connected *gorm.DB
func NewPostgresDB() *gorm.DB {

	dsn := "host=" + os.Getenv("PG_USER_HOST") +
		" user=" + os.Getenv("PG_USER_USER") +
		" password=" + os.Getenv("PG_USER_PASS") +
		" dbname=" + os.Getenv("PG_USER_NAME") +
		" port=" + os.Getenv("PG_USER_PORT") +
		" sslmode=" + os.Getenv("PG_USER_SSLMODE")

	logLevel := logger.Silent
	if os.Getenv("PG_USER_DEBUG_MODE") == "on" {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	err = AutoMigrate(db)
	if err != nil {
		panic("Migtaion Failed " + err.Error())
	}
	return db
}

func AutoMigrate(db *gorm.DB) error {

	return nil
}
