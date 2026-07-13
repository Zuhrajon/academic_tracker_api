package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) CreateSubjects(subject model.Subject) (model.Subject, error) {
	if subject.SubjectName == "" {
		return model.Subject{}, errors.New("subject_name is required")
	}

	if subject.TeacherName == "" {
		return model.Subject{}, errors.New("teacher_name is required")
	}

	if subject.TeacherUserID <= 0 {
		return model.Subject{}, errors.New("teacher_user_id is required")
	}

	teacher, err := s.repository.GetUserByID(subject.TeacherUserID)
	if err != nil {
		return model.Subject{}, errors.New("teacher user not found")
	}

	if teacher.Role != model.RoleTeacher {
		return model.Subject{}, errors.New("teacher_user_id must belong to teacher")
	}

	if subject.Semester <= 0 {
		return model.Subject{}, errors.New("Semester must be greater than zero")
	}

	if subject.Semester > 8 {
		return model.Subject{}, errors.New("Semester must be less than or equal to 8")
	}

	return s.repository.CreateSubjects(subject)
}
