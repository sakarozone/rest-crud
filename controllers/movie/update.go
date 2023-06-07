package controllers

import (
	config "learngo/restapiserver/configs"
	model "learngo/restapiserver/models"

	"github.com/gin-gonic/gin"
)

func UpdateMovie(c *gin.Context) {
	//Get id fom the url
	id := c.Param(("id"))
	//Get data from body
	var body struct {
		ID       uint
		Name     string
		Year     uint
		Director string
		Rating   uint
	}
	c.Bind(&body)

	//Find the movie to update
	var movie model.MovieTable
	config.DB.First(&movie, id)

	//Update the movie
	result := config.DB.Model(&movie).Updates(model.MovieTable{
		Name:     body.Name,
		Year:     body.Year,
		Director: body.Director,
		Rating:   body.Rating,
	})

	//return it
	var updatedMovie model.MovieTable
	result.Scan(&updatedMovie)

	c.JSON(200, gin.H{
		"movie": updatedMovie,
	})

}
