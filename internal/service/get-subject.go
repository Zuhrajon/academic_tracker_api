package service

import "academic-tracker-api/internal/model"

func (s *Service) GetSubject() ([]model.Subject, error) {
	return s.repository.GetSubject()
}
