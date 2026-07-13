package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
	"fmt"
)

func (r *Repository) GetSubjectById(id int) (model.Subject, error) {
	query := `
		SELECT id, subject_name, teacher_name, COALESCE(teacher_user_id, 0), semester
	FROM subjects
	WHERE id = $1
`
	var subject model.Subject

	err := r.db.QueryRow(query, id).Scan(
		&subject.ID,
		&subject.SubjectName,
		&subject.TeacherName,
		&subject.TeacherUserID,
		&subject.Semester,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Subject{}, fmt.Errorf("subject not found")
		}

		return model.Subject{}, fmt.Errorf("get subject by id query error: %w", err)
	}

	return subject, nil

}
