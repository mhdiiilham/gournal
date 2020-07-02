package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/handlers"
)

// Journal handlers
func Journal(r *gin.RouterGroup) {
	r.POST("/journals", handlers.PostJournal)
}
