package routers

import (
	"github.com/mhdiiilham/gournal/controllers"
	"github.com/gin-gonic/gin"
)

// AuthHandler ...
func AuthHandler(r *gin.RouterGroup) {
	r.POST("/signup", controllers.CreateUser)
	r.POST("/signin", controllers.Login)
	r.POST("/signout", controllers.Logout)
}