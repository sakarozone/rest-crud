package controllers

import (
	model "learngo/restapiserver/models"
	"learngo/restapiserver/services"

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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//return it
	c.JSON(200, gin.H{
		"movie": "Movie created successfully",
	})
}
