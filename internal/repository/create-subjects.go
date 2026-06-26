package repository

import "academic-tracker-api/internal/model"

func (r *Repository) CreateSubjects(subject model.Subject) (model.Subject, error) {
	query := `
		INSERT INTO subjects (subject_name, teacher_name, semester)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		subject.SubjectName,
		subject.TeacherName,
		subject.Semester,
	).Scan(&subject.ID)
	if err != nil {
		return model.Subject{}, err
	}

	return subject, nil
}
