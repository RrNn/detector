package models

import (
	"time"

	"gorm.io/gorm"
)

type Ping struct {
	gorm.Model
	Time   time.Time `json:"Time,omitempty"`
	Status int       `json:"Status,omitempty"`
	Error  string    `json:"Error,omitempty"`
	UrlID  uint      `json:"UrlID,omitempty"`
	Url    Url       `gorm:"references:ID"`
}
