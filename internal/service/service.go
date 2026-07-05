package service

import "academic-tracker-api/internal/model"

type Repository interface {
	GetStudents() ([]model.Student, error)
	GetStudentById(id int) (model.Student, error)
	CreateStudents(student model.Student) (model.Student, error)
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

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
