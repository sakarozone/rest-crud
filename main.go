package main

import (
	config "learngo/restapiserver/configs"
	controllers "learngo/restapiserver/controllers/movie"
	usercontrollers "learngo/restapiserver/controllers/user"
	"learngo/restapiserver/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", usercontrollers.Signup)
	r.POST("/login", usercontrollers.Login)
	r.GET("/validate", middleware.RequireAuth, usercontrollers.Validate)
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
