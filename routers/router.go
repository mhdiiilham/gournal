package routers

import (
	"github.com/mhdiiilham/gournal/middlewares"
	"github.com/gin-gonic/gin"
)

// Router ...
func Router() *gin.Engine {
	r := gin.Default()

	
	v1 := r.Group("/api/v1")
	{
		AuthHandler(v1)
	}
	r.GET("/", middlewares.Authentication(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"uid": c.MustGet("UID"),
			"email": c.MustGet("EMAIL"),
		})
	})
	return r
}
