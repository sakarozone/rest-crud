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

func GetMovie(c *gin.Context) {
	//get the movies from database and store
	var movies []model.MovieTable
	config.DB.Find(&movies)

	c.JSON(200, gin.H{
		"movie": movies,
	})
}
func GetMovieById(c *gin.Context) {
	//get the id from the url
	id := c.Param(("id"))

	var movie model.MovieTable
	config.DB.First(&movie, id)

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

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
