package main

import (
	"fmt"
	config "learngo/restapiserver/configs"
	model "learngo/restapiserver/models"
)


func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main(){
config.DB.AutoMigrate(&model.MovieTable{})
fmt.Println("Migrated successfully")
}