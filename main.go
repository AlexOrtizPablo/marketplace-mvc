package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"marketplace-mvc/controller"
	"marketplace-mvc/db"
)

func main() {
	router := gin.Default()
	err := db.ExecuteMigrations()
	if err != nil {
		log.Fatal("Cannot execute db migrations")
		return
	}
	controller.ConfigureLayers(router)
	router.Run(":8000")
}
