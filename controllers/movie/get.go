package controllers

import (
	"learngo/restapiserver/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovie(c *gin.Context) {
	err, movies := services.ListMovies()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"movie": movies,
	})

}
