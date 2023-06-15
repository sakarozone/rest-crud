package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Deletes a movie with the specified ID
// @Tags Movies
// @Accept       json
// @Produce      json
// @Param id path int true "Movie ID"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Router /movies/{id} [delete]

func (h *Handler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid movie ID",
		})
		return
	}
	err = h.Service.DeleteMovie(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete the movie",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Movie deleted successfully",
	})
}
