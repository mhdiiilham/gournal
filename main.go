package main

import (
	"github.com/mhdiiilham/gournal/db"
	"github.com/mhdiiilham/gournal/models"
	"github.com/mhdiiilham/gournal/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	db.DB()
	r := routers.Router()
	migrate()
	log.Info("SERVER IS RUNNING")
	r.Run(":8000")
}

// Migrate ...
func migrate() {
	log.Info("Start migration")
	db.DB().AutoMigrate(models.Admin{}, models.Journal{}, models.Image{})
	log.Info("Migration done!")
}
