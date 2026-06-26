package service

import "fmt"

func (s *Service) DeleteStudent(studentId int) error {
	if studentId <= 0 {
		return fmt.Errorf("invalid students id")
	}

	err := s.repository.DeleteStudent(studentId)
	if err != nil {
		return fmt.Errorf("delete student error: %w", err)
	}

	return nil
}
