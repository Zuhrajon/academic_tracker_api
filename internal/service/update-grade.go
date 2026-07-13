package service

import (
	"academic-tracker-api/internal/model"
	"errors"
	"fmt"
)

func (s *Service) UpdateGrade(gradeID int, grade model.Grade, actorUserID int, actorRole string) (model.Grade, error) {
	if gradeID <= 0 {
		return model.Grade{}, errors.New("grade_id must be greater than zero")
	}

	if grade.StudentId <= 0 {
		return model.Grade{}, errors.New("student_id must be greater than zero")
	}

	if grade.SubjectId <= 0 {
		return model.Grade{}, errors.New("subject_id must be greater than zero")
	}

	if grade.GradeDate == "" {
		return model.Grade{}, errors.New("grade_date must not be empty")
	}

	if grade.Grade <= 0 || grade.Grade > 10 {
		return model.Grade{}, errors.New("grade must be between 1 and 10")
	}

	currentGrade, err := s.repository.GetGradeByID(gradeID)
	if err != nil {
		return model.Grade{}, fmt.Errorf("get grade error: %w", err)
	}

	if err = s.ensureTeacherOwnsSubject(actorUserID, actorRole, currentGrade.SubjectId); err != nil {
		return model.Grade{}, err
	}

	if err = s.ensureTeacherOwnsSubject(actorUserID, actorRole, grade.SubjectId); err != nil {
		return model.Grade{}, err
	}

	updateGrade, err := s.repository.UpdateGrade(gradeID, grade)
	if err != nil {
		return model.Grade{}, fmt.Errorf("update grade error: %w", err)
	}

	return updateGrade, nil
}
