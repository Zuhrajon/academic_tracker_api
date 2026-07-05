package service

import "academic-tracker-api/internal/model"

func (s *Service) GetStudents() ([]model.Student, error) {
	return s.repository.GetStudents()
}
