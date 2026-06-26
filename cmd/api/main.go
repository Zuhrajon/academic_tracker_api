package main

import (
	"academic-tracker-api/internal/handler"
	"academic-tracker-api/internal/repository"
	"academic-tracker-api/internal/service"
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const defaultPostgresDSN = "postgresql://academic-tracker-user:academic-tracker-password@localhost:5432/academic-tracker-db?sslmode=disable"

func main() {
	router := gin.Default()

	postgresDSN := os.Getenv("POSTGRES_DSN")
	if postgresDSN == "" {
		postgresDSN = defaultPostgresDSN
	}

	db, err := sql.Open("pgx", postgresDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	h := handler.NewHandler(services)
	h.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
