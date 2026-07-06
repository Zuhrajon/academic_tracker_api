package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) CreateGrades(grades model.Grade) (model.Grade, error) {
	query := `
		INSERT INTO grades(student_id, subject_id, grade, grade_date, comment)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
`
	err := r.db.QueryRow(
		query,
		grades.StudentId,
		grades.SubjectId,
		grades.Grade,
		grades.GradeDate,
		grades.Comment,
	).Scan(&grades.ID)
	if err != nil {
		return model.Grade{}, fmt.Errorf("create grades query error: %w", err)
	}

	return grades, nil

}
