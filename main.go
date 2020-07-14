package main

import (
	"github.com/mhdiiilham/gournal/models"
	"github.com/mhdiiilham/gournal/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	models.ConnectDatabase()
	r := routers.Router()
	log.Info("SERVER IS RUNNING")
	r.Run(":8000")
}
