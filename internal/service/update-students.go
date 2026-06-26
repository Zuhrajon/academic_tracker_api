package service

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (s *Service) UpdateStudents(studentID int, student model.Student) (model.Student, error) {
	if studentID <= 0 {
		return model.Student{}, fmt.Errorf("invalid student id")
	}

	if student.FirstName == "" {
		return model.Student{}, fmt.Errorf("first_name is required")
	}

	if student.LastName == "" {
		return model.Student{}, fmt.Errorf("last_name is required")
	}

	if student.GroupName == "" {
		return model.Student{}, fmt.Errorf("group_name is required")
	}

	if student.Course <= 0 {
		return model.Student{}, fmt.Errorf("course must be greater than zero")
	}

	if student.Course > 6 {
		return model.Student{}, fmt.Errorf("course must be less than or equal to 6")
	}

	updatedStudent, err := s.repository.UpdateStudents(studentID, student)
	if err != nil {
		return model.Student{}, fmt.Errorf("update student error: %w", err)
	}

	return updatedStudent, nil
}
