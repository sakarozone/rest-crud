package controllers

import (
	"fmt"
	"learngo/restapiserver/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteMovie(c *gin.Context) {
	//get the id from the url
	id := c.Param(("id"))

	//Delete the movie
	// config.DB.Delete(&model.MovieTable{}, id)
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = services.DeleteMovie(num)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//return the status
	c.JSON(200, gin.H{
		"status": "Deleted successfully",
	})
}
