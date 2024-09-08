package config

import (
	"blogging-platform/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var (
		DB_HOST     = os.Getenv("DB_HOST")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_USERNAME = os.Getenv("DB_USERNAME")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		// Log to see changes
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	// migrate model
	err = DB.Debug().AutoMigrate(&model.User{}, &model.BlogPost{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return DB, nil
}
