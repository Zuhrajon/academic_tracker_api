package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) GetAttendanceByStudentID(studentID int) ([]model.Attendance, error) {
	if studentID <= 0 {
		return nil, errors.New("student_id must not be zero ")
	}

	_, err := s.repository.GetStudentById(studentID)
	if err != nil {
		return nil, err
	}

	getAttendanceById, err := s.repository.GetAttendanceByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	return getAttendanceById, nil
}
