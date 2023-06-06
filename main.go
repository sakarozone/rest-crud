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

func main() {
	r := gin.Default()
	r.POST("/create", controllers.CreateMovie)
	r.GET("/get", controllers.GetMovie)
	r.GET("/get/:id", controllers.GetMovieById)
	r.PUT("/update/:id", controllers.UpdateMovie)
	r.DELETE("/delete/:id", controllers.DeleteMovie)
	r.Run()
}
