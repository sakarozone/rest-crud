package controllers

import (
	"fmt"
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/services"
	"net/http"
	"strconv"

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

	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	movie := model.MovieTable{
		ID:       body.ID,
		Name:     body.Name,
		Year:     body.Year,
		Director: body.Director,
		Rating:   body.Rating,
	}

	err = services.UpdateMovie(num, movie)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"movie": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})

}
