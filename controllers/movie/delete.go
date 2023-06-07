package controllers

import (
	config "learngo/restapiserver/configs"
	model "learngo/restapiserver/models"

	"github.com/gin-gonic/gin"
)

func DeleteMovie(c *gin.Context) {
	//get the id from the url
	id := c.Param(("id"))

	//Delete the movie
	config.DB.Delete(&model.MovieTable{}, id)

	//return the status
	c.JSON(200, gin.H{
		"status": "Deleted successfully",
	})
}
