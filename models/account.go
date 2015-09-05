package models

type Account struct {
    ID uint `gorm:"primary_key"`
    Username string
    Password string
}
