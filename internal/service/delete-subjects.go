package service

import "fmt"

func (s *Service) DeleteSubjects(subjectId int) error {
	if subjectId <= 0 {
		return fmt.Errorf("invalid subject id")
	}

	err := s.repository.DeleteSubjects(subjectId)
	if err != nil {
		return fmt.Errorf("delete subject error: %w", err)
	}

	return nil
}
