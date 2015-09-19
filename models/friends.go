package models

import (
	"time"
)

type Follow struct {
	ID        uint `gorm:"primary_key"`
	UserId    uint
	FollowId  uint
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt *time.Time
}
