package heron

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)


func GetDatabase() (gorm.DB){
    db, err := gorm.Open("mysql", Config.DatabaseURI)
    if err != nil {
        panic(err)
    }
    return db
}
