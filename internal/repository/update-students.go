package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) UpdateStudents(studentId int, student model.Student) (model.Student, error) {
	query := `UPDATE students
	SET first_name = $1,
		last_name = $2,
		group_name = $3,
		course = $4,
		updated_at = NOW()
	WHERE id = $5
	RETURNING id, first_name, last_name, group_name, course
`

	err := r.db.QueryRow(
		query,
		student.FirstName,
		student.LastName,
		student.GroupName,
		student.Course,
		studentId,
	).Scan(
		&student.ID,
		&student.FirstName,
		&student.LastName,
		&student.GroupName,
		&student.Course,
	)

	if err != nil {
		return model.Student{}, fmt.Errorf("ошибка, студент с таким id не найден: %w", err)
	}

	return student, nil
}
