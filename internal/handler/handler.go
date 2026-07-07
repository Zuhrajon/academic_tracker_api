package handler

import (
	"academic-tracker-api/internal/middleware"
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Register(request model.RegisterRequest) (model.User, error)
	Login(request model.LoginRequest) (model.LoginResponse, error)

	CreateStudents(student model.Student) (model.Student, error)
	GetStudents() ([]model.Student, error)
	GetStudentById(id int) (model.Student, error)
	UpdateStudents(studentID int, student model.Student) (model.Student, error)
	DeleteStudent(studentId int) error

	CreateSubjects(subject model.Subject) (model.Subject, error)
	GetSubject() ([]model.Subject, error)
	GetSubjectById(id int) (model.Subject, error)
	UpdateSubject(subjectId int, subject model.Subject) (model.Subject, error)
	DeleteSubjects(subjectId int) error

	CreateAttendance(attendance model.Attendance) (model.Attendance, error)
	GetAttendanceByStudentID(studentID int) ([]model.Attendance, error)
	UpdateAttendance(attendanceID int, attendance model.Attendance) (model.Attendance, error)
	DeleteAttendance(attendanceId int) error

	CreateGrades(grades model.Grade) (model.Grade, error)
	GetGradesByStudentID(studentID int) ([]model.Grade, error)
	UpdateGrade(gradeID int, grade model.Grade) (model.Grade, error)
	DeleteGrade(gradeID int) error
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
	router.POST("/auth/register", h.Register)
	router.POST("/auth/login", h.Login)
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	authorized.GET("/students", middleware.RoleMiddleware(model.RoleAdmin, model.RoleTeacher), h.GetStudents)
	authorized.GET("/students/:id", middleware.RoleMiddleware(model.RoleAdmin, model.RoleTeacher, model.RoleStudent), h.GetStudentById)
	authorized.POST("/students", middleware.RoleMiddleware(model.RoleAdmin), h.CreateStudents)
	authorized.PUT("/students/:id", middleware.RoleMiddleware(model.RoleAdmin), h.UpdateStudents)
	authorized.DELETE("/students/:id", middleware.RoleMiddleware(model.RoleAdmin), h.DeleteStudent)

	authorized.POST("/subjects", middleware.RoleMiddleware(model.RoleAdmin), h.CreateSubjects)
	authorized.GET("/subjects", middleware.RoleMiddleware(model.RoleAdmin, model.RoleTeacher, model.RoleStudent), h.GetSubject)
	authorized.GET("/subjects/:id", middleware.RoleMiddleware(model.RoleAdmin, model.RoleTeacher, model.RoleStudent), h.GetSubjectById)
	authorized.PUT("/subjects/:id", middleware.RoleMiddleware(model.RoleAdmin), h.UpdateSubject)
	authorized.DELETE("/subjects/:id", middleware.RoleMiddleware(model.RoleAdmin), h.DeleteSubjects)

	authorized.POST("/attendance", middleware.RoleMiddleware(model.RoleTeacher), h.CreateAttendance)
	authorized.GET("/students/:id/attendance", middleware.RoleMiddleware(model.RoleAdmin, model.RoleTeacher, model.RoleStudent), h.GetAttendanceByStudentID)
	authorized.PUT("/attendance/:id", middleware.RoleMiddleware(model.RoleTeacher), h.UpdateAttendance)
	authorized.DELETE("/attendance/:id", middleware.RoleMiddleware(model.RoleTeacher), h.DeleteAttendance)

	authorized.POST("/grades", middleware.RoleMiddleware(model.RoleTeacher), h.CreateGrades)
	authorized.GET("/students/:id/grades", middleware.RoleMiddleware(model.RoleAdmin, model.RoleTeacher, model.RoleStudent), h.GetGradesByStudentID)
	authorized.PUT("/grades/:id", middleware.RoleMiddleware(model.RoleTeacher), h.UpdateGrade)
	authorized.DELETE("/grades/:id", middleware.RoleMiddleware(model.RoleTeacher), h.DeleteGrade)
}
