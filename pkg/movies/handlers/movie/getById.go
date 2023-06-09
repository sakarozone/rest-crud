package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMovieById godoc
// @Summary Get a movie by ID
// @Description Retrieves a movie by its ID
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Router /movies/{id} [get]
func (h *Handler) GetMovieById(c *gin.Context) {
	id := c.Param("id")

	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	err, movie := h.Service.ListOneMovie(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve the movie",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}
