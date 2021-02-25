package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Link   string `gorm:"type:text" json:"Link"`
	UserID uint
	User   User `gorm:"references:ID"`
}
