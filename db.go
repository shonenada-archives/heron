package heron

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shonenada/heron/models"
)

func GetDatabase() gorm.DB {
	db, err := gorm.Open("mysql", Config.DatabaseURI)
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate() {
	db := GetDatabase()
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Account{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Event{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Follow{})
	db.Model(&models.Event{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Follow{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Follow{}).AddForeignKey("follow_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&models.Account{}, &models.Follow{}, &models.Event{})
}
