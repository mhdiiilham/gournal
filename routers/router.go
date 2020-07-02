package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/middlewares"
)

// Router ...
func Router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		AuthHandler(v1)
		// Authenticated routes
		v1.Use(middlewares.Authentication())
		UploadImage(v1)
		Journal(v1)
	}
	return r
}
