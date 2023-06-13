package routes

import (
	handlers "learngo/restapiserver/pkg/movies/handlers/movie"
	"learngo/restapiserver/pkg/movies/middleware"

	"github.com/gin-gonic/gin"
)

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
		InitAuthRoutes(baseRouteAuth, baseRoute)

	}
}
