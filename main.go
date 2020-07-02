package main

import (
	"github.com/mhdiiilham/gournal/db"
	"github.com/mhdiiilham/gournal/models"
	"github.com/mhdiiilham/gournal/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	db.DB()
	r := routers.Routaer()
	log.Info("SERVER IS RUNNING")
	r.Run(":8000")
}
