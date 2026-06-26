package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) CreateStudents(student model.Student) (model.Student, error) {
	if student.FirstName == "" {
		return model.Student{}, errors.New("first_name is required")
	}

	if student.LastName == "" {
		return model.Student{}, errors.New("last_name is required")
	}

	if student.GroupName == "" {
		return model.Student{}, errors.New("group_name is required")
	}

	if student.Course <= 0 {
		return model.Student{}, errors.New("course must be greater than zero")
	}

	if student.Course > 6 {
		return model.Student{}, errors.New("course must be less than or equal to 6")
	}

	return s.repository.CreateStudents(student)
}
