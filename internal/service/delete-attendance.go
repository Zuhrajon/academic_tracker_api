package service

import "fmt"

func (s *Service) DeleteAttendance(attendanceId int) error {
	if attendanceId <= 0 {
		return fmt.Errorf("invalid attendance id")
	}

	err := s.repository.DeleteAttendance(attendanceId)
	if err != nil {
		return fmt.Errorf("delete attendance error: %w", err)
	}

	return nil
}
