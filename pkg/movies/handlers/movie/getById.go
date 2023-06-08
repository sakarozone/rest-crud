package controllers

import (
	"fmt"
	"learngo/restapiserver/pkg/movies/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovieById(c *gin.Context) {
	//get the id from the url
	id := c.Param(("id"))

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
	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}
