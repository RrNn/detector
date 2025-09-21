package models

import "gorm.io/gorm"

// Url exported ...
type Url struct {
	gorm.Model
	Link   string `gorm:"type:text" json:"Link,omitempty"`
	UserID uint
	User   User `gorm:"references:ID"`
	Pings  []Ping
}
