package service

import "fmt"

func (s *Service) DeleteAttendance(attendanceId int, actorUserID int, actorRole string) error {
	if attendanceId <= 0 {
		return fmt.Errorf("invalid attendance id")
	}

	attendance, err := s.repository.GetAttendanceByID(attendanceId)
	if err != nil {
		return fmt.Errorf("get attendance error: %w", err)
	}

	if err = s.ensureTeacherOwnsSubject(actorUserID, actorRole, attendance.SubjectId); err != nil {
		return err
	}

	err = s.repository.DeleteAttendance(attendanceId)
	if err != nil {
		return fmt.Errorf("delete attendance error: %w", err)
	}

	return nil
}
