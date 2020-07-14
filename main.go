package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/mhdiiilham/gournal/models"
	"github.com/mhdiiilham/gournal/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	models.ConnectDatabase()
	r := routers.Router()
	r.Use(cors.Default())
	log.Info("SERVER IS RUNNING")
	r.Run(os.Getenv("GO_PORT"))
}
