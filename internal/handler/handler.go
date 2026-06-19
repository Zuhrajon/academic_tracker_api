package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetStudents() []model.Student
	GetStudentById(id int) (model.Student, bool)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	router.GET("/students", h.GetStudents)
	router.GET("/student/:id", h.GetStudentById)
}
