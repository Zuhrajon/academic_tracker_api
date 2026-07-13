package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) CreateGrades(grades model.Grade) (model.Grade, error) {
	return s.CreateGradesForUser(grades, 0, "")
}

func (s *Service) CreateGradesForUser(grades model.Grade, actorUserID int, actorRole string) (model.Grade, error) {
	if grades.StudentId <= 0 {
		return model.Grade{}, errors.New("student_id must not be less than or equal to zero")
	}

	if grades.SubjectId <= 0 {
		return model.Grade{}, errors.New("subject_id must not be less than or equal to zero")
	}

	if grades.Grade <= 0 || grades.Grade > 10 {
		return model.Grade{}, errors.New("grade must not be less to zero and more than ten ")
	}

	if grades.GradeDate == "" {
		return model.Grade{}, errors.New("grade_date must not be empty ")
	}

	if err := s.ensureTeacherOwnsSubject(actorUserID, actorRole, grades.SubjectId); err != nil {
		return model.Grade{}, err
	}

	return s.repository.CreateGrades(grades)

}
