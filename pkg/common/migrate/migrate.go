package main

import (
	"fmt"
	"learngo/restapiserver/pkg/common/db"
	model "learngo/restapiserver/pkg/models"
	config "learngo/restapiserver/pkg/movies/configs"
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
