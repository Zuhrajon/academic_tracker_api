package service

import (
	"academic-tracker-api/internal/model"
	"errors"
)

func (s *Service) CreateAttendance(attendance model.Attendance) (model.Attendance, error) {
	if attendance.StudentId <= 0 {
		return model.Attendance{}, errors.New("student_id must not be less than or equal to zero")
	}

	if attendance.SubjectId <= 0 {
		return model.Attendance{}, errors.New("subject_id must not be less than or equal to zero")
	}

	if attendance.LessonDate == "" {
		return model.Attendance{}, errors.New("lesson_date must not be empty")
	}

	if attendance.Status == "" {
		return model.Attendance{}, errors.New("status must not be empty")
	}

	if attendance.Status != "present" && attendance.Status != "absent" && attendance.Status != "late" {
		return model.Attendance{}, errors.New("Status must be present/absent or late ")
	}

	return s.repository.CreateAttendance(attendance)

}
