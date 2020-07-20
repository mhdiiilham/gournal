package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mhdiiilham/gournal/helpers"
	log "github.com/sirupsen/logrus"
)

// DB instance
var DB *gorm.DB

// ConnectDatabase ...
func ConnectDatabase() {
	host := helpers.Getenv("MYSQL_HOST")
	user := helpers.Getenv("MYSQL_USER")
	pass := helpers.Getenv("MYSQL_USER_PASSWORD")
	name := "journal"
	log.Info(host, user, pass, "halloo")

	uri := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, name,
	)

	database, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Fatal("CANNOT CONNECT TO DATABASE :(")
	}
	DB = database
}