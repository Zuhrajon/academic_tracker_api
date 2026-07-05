package service

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (s *Service) GetSubjectById(id int) (model.Subject, error) {
	if id <= 0 {
		return model.Subject{}, fmt.Errorf("invalid subject id")
	}

	return s.repository.GetSubjectById(id)
}
