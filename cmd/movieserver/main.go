package main

import (
	"learngo/restapiserver/pkg/common/db"
	config "learngo/restapiserver/pkg/movies/configs"
	controllers "learngo/restapiserver/pkg/movies/handlers/movie"
	usercontrollers "learngo/restapiserver/pkg/movies/handlers/user"
	"learngo/restapiserver/pkg/movies/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	db.ConnectToDB()
}

func InitRoutes(adminRoutes *gin.RouterGroup) {
	movieRoute := adminRoutes.Group("/movies")
	baseRoute := adminRoutes.Group("")
	baseRouteAuth := adminRoutes.Group("")
	movieRoute.Use(middleware.RequireAuth)
	baseRouteAuth.Use(middleware.RequireAuth)
	{
		movieRoute.POST("",
			controllers.CreateMovie,
		)
		movieRoute.GET("",
			controllers.GetMovie,
		)
		movieRoute.GET("/:id",
			controllers.GetMovieById,
		)
		movieRoute.DELETE("/:id",
			controllers.DeleteMovie,
		)
		movieRoute.PUT("/:id",
			controllers.UpdateMovie,
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
	r := gin.Default()
	adminRoutes := r.Group("")
	InitRoutes(adminRoutes)
	r.Run()
}

// {
//     "Email":"sakar@test.com",
//     "Password":"testpass"
// }
