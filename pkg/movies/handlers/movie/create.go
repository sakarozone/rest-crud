package controllers

import (
	// model "learngo/restapiserver/pkg/movies/models"
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	//Get data from request body
	var body struct {
		ID       uint
		Name     string
		Year     uint
		Director string
		Rating   uint
	}
	c.Bind(&body)
	movie := model.MovieTable{
		ID:       body.ID,
		Name:     body.Name,
		Year:     body.Year,
		Director: body.Director,
		Rating:   body.Rating,
	}

	err := services.CreateMovie(movie)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//return it
	c.JSON(http.StatusOK, gin.H{
		"movie": "Movie created successfully",
	})
}
