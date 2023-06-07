package controllers

import (
	config "learngo/restapiserver/configs"
	model "learngo/restapiserver/models"

	"github.com/gin-gonic/gin"
)

func GetMovie(c *gin.Context) {
	//get the movies from database and store
	var movies []model.MovieTable
	config.DB.Find(&movies)

	c.JSON(200, gin.H{
		"movie": movies,
	})
}
