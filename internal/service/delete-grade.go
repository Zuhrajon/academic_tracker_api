package service

import "fmt"

func (s *Service) DeleteGrade(gradeID int) error {
	if gradeID <= 0 {
		return fmt.Errorf("invalid grade_id")
	}

	err := s.repository.DeleteGrade(gradeID)
	if err != nil {
		return fmt.Errorf("delete grade error: %w", err)
	}

	return nil
}
