package repository

import "academic-tracker-api/internal/model"

func (r *Repository) CreateAttendance(attendance model.Attendance) (model.Attendance, error) {
	query := `
	INSERT INTO attendance(student_id, subject_id,lesson_date, status, comment )
	VALUES ($1, $2,$3, $4, $5)
	RETURNING id
`
	err := r.db.QueryRow(
		query,
		attendance.StudentId,
		attendance.SubjectId,
		attendance.LessonDate,
		attendance.Status,
		attendance.Comment,
	).Scan(&attendance.ID)
	if err != nil {
		return model.Attendance{}, err
	}

	return attendance, nil
}
