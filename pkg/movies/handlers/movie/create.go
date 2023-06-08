package controllers

import (
	// model "learngo/restapiserver/pkg/movies/models"
	model "learngo/restapiserver/pkg/movies/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMovie(c *gin.Context) {
	//Get data from request body
	var body struct { ///make model out of it
		ID       uint
		Name     string
		Year     uint
		Director string
		Rating   uint
	}
	c.Bind(&body) //error handling change name to req
	movie := model.MovieTable{
		ID:       body.ID,
		Name:     body.Name,
		Year:     body.Year,
		Director: body.Director,
		Rating:   body.Rating,
	}

	err := h.Service.CreateMovie(movie)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//return it
	c.JSON(http.StatusOK, gin.H{
		"movie": "Movie created successfully",
	})
}
