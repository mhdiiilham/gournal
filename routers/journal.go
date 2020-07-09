package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/handlers"
)

// Journal handlers
func Journal(r *gin.RouterGroup) {
	j := r.Group("/journals")
	{
		j.GET("/", handlers.GetJournals)
		j.POST("/", handlers.PostJournal)
		j.GET("/:id", handlers.GetOneJournal)
	}
}
