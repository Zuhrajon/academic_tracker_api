package repository

import (
	"academic-tracker-api/internal/model"
	"fmt"
)

func (r *Repository) CreateUser(user model.User) (model.User, error) {
	query := `
		INSERT INTO users (email, password_hash, role, student_id)
		VALUES ($1, $2, $3, NULLIF($4, 0))
		RETURNING id, email, password_hash, role, COALESCE(student_id, 0)
	`

	err := r.db.QueryRow(
		query,
		user.Email,
		user.PasswordHash,
		user.Role,
		user.StudentID,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.StudentID,
	)

	if err != nil {
		return model.User{}, fmt.Errorf("create user query error: %w", err)
	}

	return user, nil
}
