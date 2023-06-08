package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
