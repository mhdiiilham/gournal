package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

// DB ...
func DB() *gorm.DB {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_USER_PASSWORD")
	name := "journal"

	uri := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	db, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Fatal("CANNOT CONNECT TO DATABASE :(")
	}

	log.Info("All clear, connected to database <3 :* ")
	return db
}
