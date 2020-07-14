package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

// DB instance
var DB *gorm.DB

// ConnectDatabase ...
func ConnectDatabase() {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_USER_PASSWORD")
	name := "journal"

	uri := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	database, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Fatal("CANNOT CONNECT TO DATABASE :(")
	}
	DB = database
}