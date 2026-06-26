package service

import "academic-tracker-api/internal/model"

func (s *Service) GetStudentById(id int) (model.Student, bool) {
	return s.repository.GetStudentById(id)
}
