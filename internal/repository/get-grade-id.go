package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
	"fmt"
)

func (r *Repository) GetGradeByID(gradeID int) (model.Grade, error) {
	query := `
		SELECT id, student_id, subject_id, grade_date, grade, comment
		FROM grades
		WHERE id = $1
	`

	var grade model.Grade

	err := r.db.QueryRow(query, gradeID).Scan(
		&grade.ID,
		&grade.StudentId,
		&grade.SubjectId,
		&grade.GradeDate,
		&grade.Grade,
		&grade.Comment,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Grade{}, fmt.Errorf("grade not found")
		}

		return model.Grade{}, fmt.Errorf("get grade by id query error: %w", err)
	}

	return grade, nil
}
