package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) GetGradesByStudentID(studentID int) ([]model.Grade, error) {
	if studentID <= 0 {
		return nil, errors.New("student_id must not be zero")
	}

	_, found := s.repository.GetStudentById(studentID)
	if !found {
		return nil, errors.New("student not found")
	}

	getGradesById, err := s.repository.GetGradesByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	return getGradesById, nil
}
