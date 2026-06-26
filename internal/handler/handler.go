package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetStudents() []model.Student
	GetStudentById(id int) (model.Student, bool)
	CreateStudents(student model.Student) (model.Student, error)
	UpdateStudents(studentID int, student model.Student) (model.Student, error)
	DeleteStudent(studentId int) error
	GetSubject() []model.Subject
	CreateSubjects(subject model.Subject) (model.Subject, error)
	GetSubjectById(id int) (model.Subject, bool)
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
	router.GET("/students/:id", h.GetStudentById)
	router.POST("/students", h.CreateStudents)
	router.PUT("/students/:id", h.UpdateStudents)
	router.DELETE("/students/:id", h.DeleteStudent)
	router.GET("/subjects", h.GetSubject)
	router.POST("/subjects", h.CreateSubjects)
	router.GET("/subjects/:id", h.GetSubjectById)
}
