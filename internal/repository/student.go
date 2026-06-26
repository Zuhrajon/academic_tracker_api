package repository

import "academic-tracker-api/internal/model"

func (r *Repository) GetStudents() []model.Student {
	query := `
		SELECT id, first_name, last_name, group_name, course
		FROM students
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.Student{}
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
			return []model.Student{}
		}

		students = append(students, student)
	}

	return students
}
