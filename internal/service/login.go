package service

import (
	"academic-tracker-api/internal/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(request model.LoginRequest) (model.LoginResponse, error) {
	if request.Email == "" {
		return model.LoginResponse{}, errors.New("email is required")
	}

	if request.Password == "" {
		return model.LoginResponse{}, errors.New("password is required")
	}

	user, err := s.repository.GetUserByEmail(request.Email)
	if err != nil {
		return model.LoginResponse{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(request.Password),
	)
	if err != nil {
		return model.LoginResponse{}, errors.New("invalid email or password")
	}

	token, err := generateToken(user)
	if err != nil {
		return model.LoginResponse{}, err
	}

	return model.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}
