package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
)

func (r *Repository) GetStudentById(id int) (model.Student, bool) {
	query := `
		SELECT id, first_name, last_name, group_name, course
		FROM students
		WHERE id = $1
	`

	var student model.Student

	err := r.db.QueryRow(query, id).Scan(
		&student.ID,
		&student.FirstName,
		&student.LastName,
		&student.GroupName,
		&student.Course,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Student{}, false
		}

		return model.Student{}, false
	}

	return student, true
}
