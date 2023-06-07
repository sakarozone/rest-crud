package controllers

import (
	"fmt"
	"learngo/restapiserver/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovieById(c *gin.Context) {
	//get the id from the url
	id := c.Param(("id"))

	// var movie model.MovieTable
	// config.DB.First(&movie, id)
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err, movie := services.ListOneMovie(num)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(200, gin.H{
		"movie": movie,
	})
}
