package main

import (
	"fmt"
	"learngo/restapiserver/pkg/common/db"
	model "learngo/restapiserver/pkg/movies/models"
)

func main() {
	db, err := db.ReturnDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db.AutoMigrate(&model.MovieTable{})
	db.AutoMigrate(&model.User{})
	fmt.Println("Migrated successfully")
}
