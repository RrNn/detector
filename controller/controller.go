package controller

import (
  "gorm.io/gorm"
)

// import "github.com/jinzhu/gorm"

type Controller struct {
  DB *gorm.DB
}
