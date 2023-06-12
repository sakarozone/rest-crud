package handlers

import (
	"fmt"
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
// @Router /movies/{id} [delete]
func (h *Handler) DeleteMovie(c *gin.Context) {
	id := c.Param(("id"))

	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = h.Service.DeleteMovie(num)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//return the status
	c.JSON(http.StatusOK, gin.H{
		"status": "Deleted successfully",
	})
}
