package handlers

import (
	model "learngo/restapiserver/pkg/movies/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMovieRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}


// CreateMovieRequest godoc
// @Summary      Create a Movie
// @Description  Endpoint to create a Movie
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Param        Movies body model.MovieTable true "MovieTable"
// @Success 200 {string}  "Movie created successfully"
// @Failure      404  {object}   object "Movie could not be created"
// @Failure 401 {object} object 
// @Router       /movies [post]
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
