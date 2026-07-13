package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) GetSubject() ([]model.Subject, error) {
	query := `
		SELECT id, subject_name, teacher_name, COALESCE(teacher_user_id, 0), semester
		FROM subjects
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("get subjects query error: %w", err)
	}
	defer rows.Close()

	subjects := make([]model.Subject, 0)

	for rows.Next() {
		var subject model.Subject

		err := rows.Scan(
			&subject.ID,
			&subject.SubjectName,
			&subject.TeacherName,
			&subject.TeacherUserID,
			&subject.Semester,
		)
		if err != nil {
			return nil, fmt.Errorf("scan subject error: %w", err)
		}

		subjects = append(subjects, subject)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("subjects rows error: %w", err)
	}

	return subjects, nil
}
