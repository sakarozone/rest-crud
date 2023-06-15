package handlers

import (
	"fmt"
	model "learngo/restapiserver/pkg/movies/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateMovie godoc
// @Summary Update a movie
// @Description Updates an existing movie
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param movie body model.CreateMovieRequest true "Movie object"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Failure 404 {object} object
// @Failure 401 {object} object
// @Router /movies/{id} [put]
func (h *Handler) UpdateMovie(c *gin.Context) {
	id := c.Param(("id"))
	req := model.CreateMovieRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	movie := req.ToMovie()

	err = h.Service.UpdateMovie(num, movie)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"movie": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})

}
