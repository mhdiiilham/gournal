package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/handlers"
)

// UploadImage hadnler
func UploadImage(r *gin.RouterGroup) {
	r.POST("/images", handlers.UploadImage)
}
