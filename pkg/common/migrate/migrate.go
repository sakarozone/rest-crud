package main

import (
	"fmt"
	"learngo/restapiserver/pkg/common/db"
	config "learngo/restapiserver/pkg/movies/configs"
	model "learngo/restapiserver/pkg/movies/models"
)

func init() {
	config.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(&model.MovieTable{})
	db.DB.AutoMigrate(&model.User{})
	fmt.Println("Migrated successfully")
}
