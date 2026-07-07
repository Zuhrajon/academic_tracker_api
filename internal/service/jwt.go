package service

import (
	"academic-tracker-api/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("academic-tracker-secret")

func generateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    user.ID,
		"email":      user.Email,
		"role":       user.Role,
		"student_id": user.StudentID,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
