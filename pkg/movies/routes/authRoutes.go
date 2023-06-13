package routes

import (
	usercontrollers "learngo/restapiserver/pkg/movies/handlers/user"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(baseRouteAuth *gin.RouterGroup, baseRoute *gin.RouterGroup) {
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
