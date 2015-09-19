package models

import (
	"time"
)

type Event struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Content   string `sql:"size:1024"`
	Due       time.Time
	UserId    uint
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt *time.Time
}
