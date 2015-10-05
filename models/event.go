package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Event struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Content   string `sql:"size:1024"`
	Due       time.Time
	AccountId uint
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt *time.Time
}

func (event *Event) GetUser(db gorm.DB) Account {
	accountId := event.AccountId
	account := Account{}
	db.Where("id = ?", accountId).First(&account)
	return account
}
