package main

import (
	"learn"
	config "learngo/restapiserver/configs"
	controllers "learngo/restapiserver/controllers/movie"
	"learngo/restapiserver/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/movie", middleware.RequireAuth, controllers.CreateMovie)
	r.GET("/movie", middleware.RequireAuth, controllers.GetMovie)
	r.GET("/movie/:id", middleware.RequireAuth, controllers.GetMovieById)
	r.PUT("/movie/:id", middleware.RequireAuth, controllers.UpdateMovie)
	r.DELETE("/movie/:id", middleware.RequireAuth, controllers.DeleteMovie)
	r.Run()
}

// {
//     "Email":"sakar@test.com",
//     "Password":"testpass"
// }
