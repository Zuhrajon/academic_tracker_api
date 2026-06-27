package service

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (s *Service) UpdateSubject(subjectId int, subject model.Subject) (model.Subject, error) {
	if subjectId <= 0 {
		return model.Subject{}, fmt.Errorf("invalid subject id")
	}

	if subject.SubjectName == "" {
		return model.Subject{}, fmt.Errorf("subject_name is required")
	}

	if subject.TeacherName == "" {
		return model.Subject{}, fmt.Errorf("teacher_name is required")
	}

	if subject.Semester <= 0 {
		return model.Subject{}, fmt.Errorf("semester must be greater than zero")
	}

	if subject.Semester > 8 {
		return model.Subject{}, fmt.Errorf("semester must be less than or equal to 8")
	}

	updatedSubject, err := s.repository.UpdateSubject(subjectId, subject)
	if err != nil {
		return model.Subject{}, fmt.Errorf("update subject error: %w", err)
	}

	return updatedSubject, nil
}
