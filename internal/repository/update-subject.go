package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) UpdateSubject(subjectId int, subjects model.Subject) (model.Subject, error) {
	query := `UPDATE subjects
	SET subject_name = $1,
		teacher_name = $2,
		semester = $3,
		updated_at = NOW()
	WHERE id = $4
	RETURNING id, subject_name, teacher_name, semester
`

	err := r.db.QueryRow(
		query,
		subjects.SubjectName,
		subjects.TeacherName,
		subjects.Semester,
		subjectId,
	).Scan(
		&subjects.ID,
		&subjects.SubjectName,
		&subjects.TeacherName,
		&subjects.Semester,
	)

	if err != nil {
		return model.Subject{}, fmt.Errorf("ошибка, предмет с таким id не найден: %w", err)
	}

	return subjects, nil
}
