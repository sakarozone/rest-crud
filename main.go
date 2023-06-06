package main

import (
	config "learngo/restapiserver/configs"
	"learngo/restapiserver/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main()  {
	r := gin.Default()
	r.POST("/create", controllers.CreateMovie)
	r.Run()
}