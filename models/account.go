package models

import (
	"time"
)

type Account struct {
	ID        uint `gorm:"primary_key"`
	Username  string
	Password  string
	Name      string
	Avatar    string
	Actived   bool
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt *time.Time
}
