package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
	"fmt"
)

func (r *Repository) UpdateSubject(subjectId int, subjects model.Subject) (model.Subject, error) {
	query := `UPDATE subjects
	SET subject_name = $1,
		teacher_name = $2,
		teacher_user_id = NULLIF($3, 0),
		semester = $4,
		updated_at = NOW()
	WHERE id = $5
	RETURNING id, subject_name, teacher_name, COALESCE(teacher_user_id, 0), semester
`

	err := r.db.QueryRow(
		query,
		subjects.SubjectName,
		subjects.TeacherName,
		subjects.TeacherUserID,
		subjects.Semester,
		subjectId,
	).Scan(
		&subjects.ID,
		&subjects.SubjectName,
		&subjects.TeacherName,
		&subjects.TeacherUserID,
		&subjects.Semester,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Subject{}, fmt.Errorf("subject not found")
		}
		return model.Subject{}, fmt.Errorf("update subject query error: %w", err)
	}

	return subjects, nil
}
