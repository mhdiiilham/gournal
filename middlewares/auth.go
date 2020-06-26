package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/helpers"
	log "github.com/sirupsen/logrus"
)

// Authentication ...
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		err := helpers.TokenValidation(authHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "TOKEN NOT VALID"})
			c.Abort()
			return
		}
		
		metaData, err := helpers.ExtractTokenMetaData(authHeader)
		log.Info(metaData.ID, metaData.Email)
		c.Set("UID", metaData.ID)
		c.Next()
	}
}

