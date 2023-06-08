package handlers

import (
	model "learngo/restapiserver/pkg/movies/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMovie(c *gin.Context) {
	req := model.CreateMovieRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}

	movie := req.ToMovie()
	err = h.Service.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movie": "Movie created successfully",
	})
}
