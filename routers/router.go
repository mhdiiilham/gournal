package routers

import (
	"github.com/mhdiiilham/gournal/middlewares"
	"github.com/gin-gonic/gin"
)

// Router ...
func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Authentication())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"uid":c.MustGet("UID")})
	})
	v1 := r.Group("/api/v1")
	{
		AuthHandler(v1)
	}

	return r
}
