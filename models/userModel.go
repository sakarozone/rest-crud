package model

type User struct {
	Email    string
	Password string `gorm:"unique"`
}
