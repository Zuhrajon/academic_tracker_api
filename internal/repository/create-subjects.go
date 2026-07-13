package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) CreateSubjects(subject model.Subject) (model.Subject, error) {
	query := `
		INSERT INTO subjects (subject_name, teacher_name, teacher_user_id, semester)
		VALUES ($1, $2, NULLIF($3, 0), $4)
		RETURNING id, subject_name, teacher_name, COALESCE(teacher_user_id, 0), semester
	`

	err := r.db.QueryRow(
		query,
		subject.SubjectName,
		subject.TeacherName,
		subject.TeacherUserID,
		subject.Semester,
	).Scan(
		&subject.ID,
		&subject.SubjectName,
		&subject.TeacherName,
		&subject.TeacherUserID,
		&subject.Semester,
	)
	if err != nil {
		return model.Subject{}, fmt.Errorf("create subject query error: %w", err)
	}

	return subject, nil
}
