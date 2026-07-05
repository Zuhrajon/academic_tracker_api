package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) GetStudents() ([]model.Student, error) {
	query := `
		SELECT id, first_name, last_name, group_name, course
		FROM students
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("get students query error: %w", err)
	}
	defer rows.Close()

	students := make([]model.Student, 0)

	for rows.Next() {
		var student model.Student

		err := rows.Scan(
			&student.ID,
			&student.FirstName,
			&student.LastName,
			&student.GroupName,
			&student.Course,
		)
		if err != nil {
			return nil, fmt.Errorf("scan student error: %w", err)
		}

		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("students rows error: %w", err)
	}

	return students, nil
}
