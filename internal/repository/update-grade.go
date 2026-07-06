package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
	"fmt"
)

func (r *Repository) UpdateGrade(gradeID int, grade model.Grade) (model.Grade, error) {
	query := `
		UPDATE grades
	SET student_id = $1,
		subject_id = $2,
		grade_date = $3,
		grade = $4,
		comment = $5,
		updated_at = NOW()
	WHERE id = $6
	RETURNING id, student_id,subject_id, grade_date, grade, comment
`

	err := r.db.QueryRow(
		query,
		grade.StudentId,
		grade.SubjectId,
		grade.GradeDate,
		grade.Grade,
		grade.Comment,
		gradeID,
	).Scan(
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
		return model.Grade{}, fmt.Errorf("update grade query error: %w", err)
	}
	return grade, nil
}
