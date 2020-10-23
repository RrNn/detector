package models

import (
  "time"

  "gorm.io/gorm"
)

type Ping struct {
  gorm.Model
  Time   time.Time
  Status int
  Error  string
  UrlID  uint
  Url    Url `gorm:"references:ID"`
}
