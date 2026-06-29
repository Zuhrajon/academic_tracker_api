package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) UpdateAttendance(attendanceID int, attendance model.Attendance) (model.Attendance, error) {
	if attendanceID <= 0 {
		return model.Attendance{}, errors.New("attendance_id must be greater than zero ")
	}

	if attendance.StudentId <= 0 {
		return model.Attendance{}, errors.New("student_id must be greater than zero")
	}

	if attendance.SubjectId <= 0 {
		return model.Attendance{}, errors.New("subject_id must be greater than zero")
	}

	if attendance.LessonDate == "" {
		return model.Attendance{}, errors.New("lesson_date is required")
	}

	if attendance.Status != "present" && attendance.Status != "absent" && attendance.Status != "late" {
		return model.Attendance{}, errors.New("Status must be present/absent or late ")
	}

	return s.repository.UpdateAttendance(attendanceID, attendance)
}
