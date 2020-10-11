package models

import "gorm.io/gorm"

type Url struct {
  gorm.Model
  Link   string `gorm:"type:text" json:"Link"`
  UserID int    `gorm:"type:int;"`
  User   User   `gorm:"foreignKey:UserID"`
}
