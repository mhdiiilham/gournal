package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/helpers"
	"github.com/mhdiiilham/gournal/models"
	log "github.com/sirupsen/logrus"
)

// CreateUser ...
func CreateUser(c *gin.Context) {
	var body models.AdminSignUp

	// Validating user's input
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	hashedPassword, err := helpers.HashPassword(body.Password)
	if err != nil {
		log.Fatal("ERROR ON HASHING PASSWORD")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "INTERNAL SERVER ERROR!"})
		return
	}

	admin := models.Admin{
		Fullname:       body.Fullname,
		Email:          body.Email,
		PasswordHashed: hashedPassword,
	}

	admin.Save()

	token, err := helpers.CreateToken(admin.ID, admin.Email)
	if err != nil {
		log.Fatal("ERROR ON CREATING JWT TOKEN, err:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "INTERNAL SERVER ERROR!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created!",
		"access_token": token,
		"admin_fullname": admin.Fullname,
	})
}
