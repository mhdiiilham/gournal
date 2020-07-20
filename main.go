package main

import (

	"github.com/gin-contrib/cors"
	"github.com/mhdiiilham/gournal/models"
	"github.com/mhdiiilham/gournal/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	const PORT = ":8080"
	models.ConnectDatabase()
	r := routers.Router()
	r.Use(cors.Default())
	log.Info("SERVER IS RUNNING")
	models.DB.AutoMigrate(models.Admin{}, models.Image{}, models.Journal{})
	r.Run(PORT)
}
