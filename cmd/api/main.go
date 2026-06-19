package main

import (
	"academic-tracker-api/internal/handler"
	"academic-tracker-api/internal/repository"
	"academic-tracker-api/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	h := handler.NewHandler(services)
	h.InitRoutes(router)

	err := router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}

}
