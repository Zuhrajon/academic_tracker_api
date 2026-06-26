package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
)

func (r *Repository) GetSubjectById(id int) (model.Subject, bool) {
	query := `
		SELECT id, subject_name, teacher_name, semester
	FROM subjects
	WHERE id = $1
`
	var subject model.Subject

	err := r.db.QueryRow(query, id).Scan(
		&subject.ID,
		&subject.SubjectName,
		&subject.TeacherName,
		&subject.Semester,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Subject{}, false
		}

		return model.Subject{}, false
	}

	return subject, true

}
