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
	r.POST("/signup", controllers.Signup)
	r.POST("/createmovie", controllers.CreateMovie)
	r.GET("/listmovie", controllers.GetMovie)
	r.GET("/listmovie/:id", controllers.GetMovieById)
	r.PUT("/updatemovie/:id", controllers.UpdateMovie)
	r.DELETE("/deletemovie/:id", controllers.DeleteMovie)
	r.Run()
}

// testpass is password for sakar@test.com
