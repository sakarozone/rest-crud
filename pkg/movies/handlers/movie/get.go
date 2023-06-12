package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMovie godoc
// @Summary Get a list of movies
// @Description Retrieves a paginated list of movies
// @Tags Movies
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Page size (default: 10)"
// @Success 200 {object} object 
// @Failure 500 {object} object
// @Failure 401 {object} object 
// @Router /movies [get]
func (h *Handler) GetMovie(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1 // Set default page number if not provided or invalid
	}
	pageSizeStr := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10 // Set default page size if not provided or invalid
	}

	err, movies := h.Service.ListMovies(page, pageSize)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		c.JSON(500, gin.H{"error": "Failed to retrieve items"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movie": movies,
	})

}
