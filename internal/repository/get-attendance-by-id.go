package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
	"fmt"
)

func (r *Repository) GetAttendanceByID(attendanceID int) (model.Attendance, error) {
	query := `
		SELECT id, student_id, subject_id, lesson_date, status, comment
		FROM attendance
		WHERE id = $1
	`

	var attendance model.Attendance

	err := r.db.QueryRow(query, attendanceID).Scan(
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

		return model.Attendance{}, fmt.Errorf("get attendance by id query error: %w", err)
	}

	return attendance, nil
}
