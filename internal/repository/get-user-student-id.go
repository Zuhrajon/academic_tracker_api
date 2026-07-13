package repository

import "fmt"

func (r *Repository) UserExistsByStudentID(studentID int) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE student_id = $1
		)
	`

	var exists bool

	err := r.db.QueryRow(query, studentID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check user by student id query error: %w", err)
	}

	return exists, nil
}
