package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) GetGradesByStudentID(studentID int) ([]model.Grade, error) {
	query := `
		SELECT id, student_id, subject_id, grade_date, grade, comment
		FROM grades
		WHERE student_id = $1;
`
	rows, err := r.db.Query(query, studentID)
	if err != nil {
		return nil, fmt.Errorf("get grades by student id query error: %w", err)
	}
	defer rows.Close()

	grades := make([]model.Grade, 0)

	for rows.Next() {
		var grade model.Grade

		err = rows.Scan(
			&grade.ID,
			&grade.StudentId,
			&grade.SubjectId,
			&grade.GradeDate,
			&grade.Grade,
			&grade.Comment,
		)
		if err != nil {
			return nil, fmt.Errorf("scan grade error: %w", err)
		}

		grades = append(grades, grade)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("grades rows error: %w", err)
	}

	return grades, nil

}
