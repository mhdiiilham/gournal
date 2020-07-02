package routers

import (
	"github.com/mhdiiilham/gournal/handlers"
	"github.com/gin-gonic/gin"
)

// AuthHandler ...
func AuthHandler(r *gin.RouterGroup) {
	r.POST("/signup", handlers.CreateUser)
	r.POST("/signin", handlers.Login)
	r.POST("/signout", handlers.Logout)
}