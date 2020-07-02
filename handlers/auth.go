package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/helpers"
	"github.com/mhdiiilham/gournal/models"
	log "github.com/sirupsen/logrus"
)

type resData struct {
	AdminFullname string `json:"admin_fullname"`
}

// CreateUser ...
func CreateUser(c *gin.Context) {
	var data resData
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

	c.SetCookie("auth_token", token, 3600*24, "/", os.Getenv("APP_DOMAIN"), false, true)

	data.AdminFullname = admin.Fullname
	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
		"message": "User Created!",
		"data": data.AdminFullname,
	})
}

// Login ...
func Login(c *gin.Context) {
	var data resData
	var body models.AdminSignIn

	// Validating user's input
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	credential, err := models.First(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Email / Password is Wrong!"})
		return
	}

	token, err := helpers.CreateToken(credential.ID, credential.Email)
	if err != nil {
		log.Fatal("ERROR ON CREATING JWT TOKEN, err:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "INTERNAL SERVER ERROR!"})
		return
	}

	c.SetCookie("auth_token", token, 3600*24, "/", os.Getenv("APP_DOMAIN"), false, true)

	data.AdminFullname = credential.Fullname
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Login success",
		"data": data.AdminFullname,
	})

}

// Logout ...
// Remove token from cookies
func Logout(c *gin.Context) {
	c.SetCookie("auth_token", "", 1, "/", os.Getenv("APP_DOMAIN"), false, true)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Signout success",
		"data": "",
	})
}
