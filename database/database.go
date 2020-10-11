package database

import (
  "fmt"
  "log"

  "gorm.io/driver/postgres"

  // _ "github.com/jinzhu/gorm/dialects/postgres"

  "gorm.io/gorm"
  // "github.com/jinzhu/gorm"

  "github.com/RrNn/detector/constants"
  "github.com/RrNn/detector/models"
  "github.com/joho/godotenv"
)

var dsn string = fmt.Sprintf("user=%s dbname=%s port=%s sslmode=%s password=%s TimeZone=Africa/Kampala", constants.DBUser, constants.DBName, constants.DBPort, constants.DBSslmode, constants.DBPassword)

func Migrate() (database *gorm.DB, err error) {
  err = godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file", err)
  }

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if db != nil {
    if err = db.AutoMigrate(&models.User{}, &models.Url{}); err != nil {
      return nil, err
    }
  }

  return db, nil
}

func Connect() (db *gorm.DB) {
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    log.Fatal("failed to connect database", err)
    panic("failed to connect database")
  }

  return db
}
