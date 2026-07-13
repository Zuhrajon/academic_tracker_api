package service

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (s *Service) UpdateSubject(subjectId int, subject model.Subject) (model.Subject, error) {
	if subjectId <= 0 {
		return model.Subject{}, fmt.Errorf("invalid subject id")
	}

	if subject.SubjectName == "" {
		return model.Subject{}, fmt.Errorf("subject_name is required")
	}

	if subject.TeacherName == "" {
		return model.Subject{}, fmt.Errorf("teacher_name is required")
	}

	if subject.TeacherUserID <= 0 {
		return model.Subject{}, fmt.Errorf("teacher_user_id is required")
	}

	teacher, err := s.repository.GetUserByID(subject.TeacherUserID)
	if err != nil {
		return model.Subject{}, fmt.Errorf("teacher user not found")
	}

	if teacher.Role != model.RoleTeacher {
		return model.Subject{}, fmt.Errorf("teacher_user_id must belong to teacher")
	}

	if subject.Semester <= 0 {
		return model.Subject{}, fmt.Errorf("semester must be greater than zero")
	}

	if subject.Semester > 8 {
		return model.Subject{}, fmt.Errorf("semester must be less than or equal to 8")
	}

	updatedSubject, err := s.repository.UpdateSubject(subjectId, subject)
	if err != nil {
		return model.Subject{}, fmt.Errorf("update subject error: %w", err)
	}

	return updatedSubject, nil
}
