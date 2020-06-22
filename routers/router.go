package routers

import "github.com/gin-gonic/gin"

// Router ...
func Router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		AuthHandler(v1)
	}

	return r
}
