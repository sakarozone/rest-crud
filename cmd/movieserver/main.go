package main

import (
	config "learngo/restapiserver/configs"
	"learngo/restapiserver/pkg/common/db"
	"learngo/restapiserver/pkg/movies/routes"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	db.ReturnDB()
}

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}
	r := gin.Default()
	r.GET("/docs/any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	adminRoutes := r.Group("")
	routes.InitRoutes(adminRoutes)
	r.Run(":" + strconv.Itoa(config.Port))
}

// {
//     "Email":"sakar@test.com",
//     "Password":"testpass"
// }
