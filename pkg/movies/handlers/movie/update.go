package handlers

import (
	"fmt"
	model "learngo/restapiserver/pkg/movies/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
		return
	}
	movie := req.ToMovie()

	err = h.Service.UpdateMovie(num, movie)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"movie": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})

}
