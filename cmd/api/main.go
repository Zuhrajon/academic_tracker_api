package main

import (
	"academic-tracker-api/internal/handler"
	"academic-tracker-api/internal/repository"
	"academic-tracker-api/internal/service"
	"academic-tracker-api/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	database, err := db.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	repositories := repository.NewRepository(database)
	services := service.NewService(repositories)
	h := handler.NewHandler(services)
	h.InitRoutes(router)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
