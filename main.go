package main

import (
	"github.com/mhdiiilham/gournal/db"
	"github.com/mhdiiilham/gournal/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	const APPNAME = "journal"
	db.DB(APPNAME)
	r := routers.Router()
	log.Info("SERVER IS RUNNING")
	r.Run(":8000")
}