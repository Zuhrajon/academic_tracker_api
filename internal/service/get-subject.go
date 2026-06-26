package service

import "academic-tracker-api/internal/model"

func (s *Service) GetSubject() []model.Subject {
	return s.repository.GetSubject()
}
