package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	// _ "github.com/jinzhu/gorm/dialects/postgres"

	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"

	"detector/constants"
	"detector/models"
)

func dsnString() string {
	env := constants.GetEnvironment()
	return fmt.Sprintf(
		"user=%s dbname=%s port=%s sslmode=%s password=%s TimeZone=Africa/Kampala",
		env.DBUser, env.DBName, env.DBPort, env.DBSslmode, env.DBPassword)
}

// Migrate exported
func Migrate() (database *gorm.DB, err error) {
	var dsn string = dsnString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if db != nil {
		if err = db.AutoMigrate(&models.User{}, &models.Url{}, &models.Ping{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}

// Connect exported
func Connect() (db *gorm.DB) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var dsn string = dsnString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("failed to connect database", err)
		panic("failed to connect database")
	}

	return db
}
