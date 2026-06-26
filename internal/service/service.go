package service

import "academic-tracker-api/internal/model"

type Repository interface {
	GetStudents() []model.Student
	GetStudentById(id int) (model.Student, bool)
	CreateStudents(student model.Student) (model.Student, error)
	UpdateStudents(studentID int, student model.Student) (model.Student, error)
	DeleteStudent(studentId int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetStudents() []model.Student {
	return s.repository.GetStudents()
}
