package main

import (
	"learngo/restapiserver/pkg/common/db"
	config "learngo/restapiserver/pkg/movies/configs"
	handlers "learngo/restapiserver/pkg/movies/handlers/movie"
	usercontrollers "learngo/restapiserver/pkg/movies/handlers/user"
	"learngo/restapiserver/pkg/movies/middleware"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ReturnDB()
}

func InitRoutes(adminRoutes *gin.RouterGroup) {
	mh := handlers.New()
	movieRoute := adminRoutes.Group("/movies")
	baseRoute := adminRoutes.Group("")
	baseRouteAuth := adminRoutes.Group("")
	movieRoute.Use(middleware.RequireAuth)
	baseRouteAuth.Use(middleware.RequireAuth)

	{
		movieRoute.POST("",
			mh.CreateMovie,
		)
		movieRoute.GET("",
			mh.GetMovie,
		)
		movieRoute.GET("/:id",
			mh.GetMovieById,
		)
		movieRoute.DELETE("/:id",
			mh.DeleteMovie,
		)
		movieRoute.PUT("/:id",
			mh.UpdateMovie,
		)
		baseRoute.POST("/signup",
			usercontrollers.Signup,
		)
		baseRoute.POST("/login",
			usercontrollers.Login,
		)
		baseRouteAuth.GET("/validate",
			usercontrollers.Validate,
		)

	}
}

func main() {
	config, err := config.ReadConfig()

	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}
	r := gin.Default()
	adminRoutes := r.Group("")
	InitRoutes(adminRoutes)
	r.Run(":" + strconv.Itoa(config.Port))
}

// {
//     "Email":"sakar@test.com",
//     "Password":"testpass"
// }
