package controllers

import (
	config "learngo/restapiserver/configs"
	model "learngo/restapiserver/models"

	"github.com/gin-gonic/gin"
)

func GetMovieById(c *gin.Context) {
	//get the id from the url
	id := c.Param(("id"))

	var movie model.MovieTable
	config.DB.First(&movie, id)

	c.JSON(200, gin.H{
		"movie": movie,
	})
}
