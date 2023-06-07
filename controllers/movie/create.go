package controllers

import (
	config "learngo/restapiserver/configs"
	model "learngo/restapiserver/models"

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

	//Create an entry in database
	movie := model.MovieTable{ID: body.ID, Name: body.Name, Year: body.Year, Director: body.Director, Rating: body.Rating}
	result := config.DB.Create(&movie)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	var createdMovie model.MovieTable
	result.Scan(&createdMovie)

	//return it
	c.JSON(200, gin.H{
		"movie": createdMovie,
	})
}
