package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) CreateSubjects(subject model.Subject) (model.Subject, error) {
	if subject.SubjectName == "" {
		return model.Subject{}, errors.New("subject_name is required")
	}

	if subject.TeacherName == "" {
		return model.Subject{}, errors.New("teacher_name is required")
	}

	if subject.Semester <= 0 {
		return model.Subject{}, errors.New("Semester must be greater than zero")
	}

	if subject.Semester > 8 {
		return model.Subject{}, errors.New("Semester must be less than or equal to 8")
	}

	return s.repository.CreateSubjects(subject)
}
