package repository

import (
	"academic-tracker-api/internal/model"
	"database/sql"
	"fmt"
)

func (r *Repository) GetUserByID(id int) (model.User, error) {
	query := `
		SELECT id, email, password_hash, role, COALESCE(student_id, 0)
		FROM users
		WHERE id = $1
	`

	var user model.User

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.StudentID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("user not found")
		}

		return model.User{}, fmt.Errorf("get user by id query error: %w", err)
	}

	return user, nil
}
