package service

import (
	"academic-tracker-api/internal/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(request model.RegisterRequest) (model.User, error) {
	if request.Email == "" {
		return model.User{}, errors.New("email is required")
	}

	if request.Password == "" {
		return model.User{}, errors.New("password is required")
	}

	if len(request.Password) < 8 {
		return model.User{}, errors.New("password must be at least 8 characters")
	}

	if request.Role != model.RoleAdmin && request.Role != model.RoleTeacher && request.Role != model.RoleStudent {
		return model.User{}, errors.New("role must be admin, teacher or student")
	}

	if request.Role == model.RoleStudent && request.StudentID <= 0 {
		return model.User{}, errors.New("student_id is required for student role")
	}

	if request.Role != model.RoleStudent {
		request.StudentID = 0
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Email:        request.Email,
		PasswordHash: string(passwordHash),
		Role:         request.Role,
		StudentID:    request.StudentID,
	}

	createdUser, err := s.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}
