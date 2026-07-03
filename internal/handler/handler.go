package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateStudents(student model.Student) (model.Student, error)
	GetStudents() []model.Student
	GetStudentById(id int) (model.Student, bool)
	UpdateStudents(studentID int, student model.Student) (model.Student, error)
	DeleteStudent(studentId int) error

	CreateSubjects(subject model.Subject) (model.Subject, error)
	GetSubject() []model.Subject
	GetSubjectById(id int) (model.Subject, bool)
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
	router.GET("/students", h.GetStudents)
	router.GET("/students/:id", h.GetStudentById)
	router.POST("/students", h.CreateStudents)
	router.PUT("/students/:id", h.UpdateStudents)
	router.DELETE("/students/:id", h.DeleteStudent)

	router.POST("/subjects", h.CreateSubjects)
	router.GET("/subjects", h.GetSubject)
	router.GET("/subjects/:id", h.GetSubjectById)
	router.PUT("/subjects/:id", h.UpdateSubject)
	router.DELETE("/subjects/:id", h.DeleteSubjects)

	router.POST("/attendance", h.CreateAttendance)
	router.GET("/students/:id/attendance", h.GetAttendanceByStudentID)
	router.PUT("/attendance/:id", h.UpdateAttendance)
	router.DELETE("/attendance/:id", h.DeleteAttendance)

	router.POST("/grades", h.CreateGrades)
	router.GET("/students/:id/grades", h.GetGradesByStudentID)
	router.PUT("/grades/:id", h.UpdateGrade)
	router.DELETE("/grades/:id", h.DeleteGrade)

}
