package db

import (
	"fmt"
	"os"

	"github.com/mhdiiilham/gournal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

// DB ...
func DB(name string) *gorm.DB {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_USER_PASSWORD")

	uri := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	db, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Fatal("CANNOT CONNECT TO DATABASE!")
	}
	db.AutoMigrate(models.Admin{}, models.Journal{})
	log.Info("All clear, connected to database!")
	return db
}