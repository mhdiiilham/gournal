package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler ...
func AuthHandler(r *gin.RouterGroup) {
	r.GET("/its-works", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "IT'S WORKS!"})
	})
}