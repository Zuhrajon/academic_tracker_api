package service

import "fmt"

func (s *Service) DeleteStudent(studentId int) error {
	err := s.repository.DeleteStudent(studentId)
	if err != nil {
		return fmt.Errorf("delete student error: %w", err)
	}

	return nil
}
