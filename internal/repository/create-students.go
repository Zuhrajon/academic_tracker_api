package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) CreateStudents(student model.Student) (model.Student, error) {
	query := `
		INSERT INTO students (first_name, last_name, group_name, course)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		student.FirstName,
		student.LastName,
		student.GroupName,
		student.Course,
	).Scan(&student.ID)
	if err != nil {
		return model.Student{}, fmt.Errorf("create student query error: %w", err)
	}

	return student, nil
}
