package repository

import "academic-tracker-api/internal/model"

func (r *Repository) GetSubject() []model.Subject {
	query := `
		SELECT id, subject_name, teacher_name, semester
		FROM subjects
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.Subject{}
	}
	defer rows.Close()

	subjects := make([]model.Subject, 0)

	for rows.Next() {
		var subject model.Subject

		err := rows.Scan(
			&subject.ID,
			&subject.SubjectName,
			&subject.TeacherName,
			&subject.Semester,
		)
		if err != nil {
			return []model.Subject{}
		}

		subjects = append(subjects, subject)
	}

	return subjects
}
