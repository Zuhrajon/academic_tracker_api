package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
	"fmt"
)

func (r *Repository) UpdateAttendance(attendanceID int, attendance model.Attendance) (model.Attendance, error) {
	query := `
		UPDATE attendance
	SET student_id = $1,
		subject_id = $2,
		lesson_date = $3,
		status = $4,
		comment = $5,
		updated_at = NOW()
	WHERE id = $6
	RETURNING id, student_id, subject_id, lesson_date, status, comment
`
	err := r.db.QueryRow(
		query,
		attendance.StudentId,
		attendance.SubjectId,
		attendance.LessonDate,
		attendance.Status,
		attendance.Comment,
		attendanceID,
	).Scan(
		&attendance.ID,
		&attendance.StudentId,
		&attendance.SubjectId,
		&attendance.LessonDate,
		&attendance.Status,
		&attendance.Comment,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Attendance{}, fmt.Errorf("attendance not found")
		}

		return model.Attendance{}, fmt.Errorf("update attendance query error: %w", err)
	}

	return attendance, nil
}
