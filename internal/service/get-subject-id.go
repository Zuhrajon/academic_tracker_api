package service

import (
	"academic-tracker-api/internal/model"
)

func (s *Service) GetSubjectById(id int) (model.Subject, bool) {
	return s.repository.GetSubjectById(id)
}
