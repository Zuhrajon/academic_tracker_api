package service

import "fmt"

func (s *Service) DeleteGrade(gradeID int, actorUserID int, actorRole string) error {
	if gradeID <= 0 {
		return fmt.Errorf("invalid grade_id")
	}

	grade, err := s.repository.GetGradeByID(gradeID)
	if err != nil {
		return fmt.Errorf("get grade error: %w", err)
	}

	if err = s.ensureTeacherOwnsSubject(actorUserID, actorRole, grade.SubjectId); err != nil {
		return err
	}

	err = s.repository.DeleteGrade(gradeID)
	if err != nil {
		return fmt.Errorf("delete grade error: %w", err)
	}

	return nil
}
