package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
