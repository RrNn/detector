package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:text" json:"Name,omitempty"`
	Email    string `gorm:"type:text; not null; unique;" json:"Email,omitempty"`
	Password string `gorm:"type:text; not null;" json:"Password,omitempty"`
	Address  string `gorm:"type:text; not null;" json:"Address,omitempty"`
	Contact  string `gorm:"type:text; not null; unique;" json:"Contact,omitempty"`
}

type ReturnUser struct {
	ID      uint   `json:"ID"`
	Name    string `json:"Name"`
	Email   string `json:"Email"`
	Address string `json:"Address"`
	Contact string `json:"Contact"`
}

func GetUser(u *User) ReturnUser {
	return ReturnUser{
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Address: u.Address,
		Contact: u.Contact,
	}
}

type LoginUserStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
