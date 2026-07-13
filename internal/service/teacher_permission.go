package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) ensureTeacherOwnsSubject(actorUserID int, actorRole string, subjectID int) error {
	if actorRole != model.RoleTeacher {
		return nil
	}

	if actorUserID <= 0 {
		return errors.New("invalid auth user")
	}

	subject, err := s.repository.GetSubjectById(subjectID)
	if err != nil {
		return err
	}

	if subject.TeacherUserID != actorUserID {
		return errors.New("access denied")
	}

	return nil
}
