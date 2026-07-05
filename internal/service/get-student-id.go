package service

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (s *Service) GetStudentById(id int) (model.Student, error) {
	if id <= 0 {
		return model.Student{}, fmt.Errorf("invalid student id")
	}

	return s.repository.GetStudentById(id)
}
