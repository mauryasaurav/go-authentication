package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
	FirstName string `json:"first_name" binding:"required"`
}

func TestHere(c *gin.Context) {
	created := new(Test)
	if err := c.Bind(&created); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": created})
}
