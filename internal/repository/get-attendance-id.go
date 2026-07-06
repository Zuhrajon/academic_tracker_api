package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) GetAttendanceByStudentID(studentID int) ([]model.Attendance, error) {
	query := `
		SELECT id, student_id, subject_id, lesson_date, status, comment
		FROM attendance
		WHERE student_id = $1;
`
	rows, err := r.db.Query(query, studentID)
	if err != nil {
		return nil, fmt.Errorf("get attendance by student id query error: %w", err)
	}
	defer rows.Close()

	attendances := make([]model.Attendance, 0)

	for rows.Next() {
		var attendance model.Attendance

		err = rows.Scan(
			&attendance.ID,
			&attendance.StudentId,
			&attendance.SubjectId,
			&attendance.LessonDate,
			&attendance.Status,
			&attendance.Comment,
		)
		if err != nil {
			return nil, fmt.Errorf("scan attendance error: %w", err)
		}

		attendances = append(attendances, attendance)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("attendance rows error: %w", err)
	}

	return attendances, nil
}
