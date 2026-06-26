package repository

import "fmt"

func (r *Repository) DeleteStudent(studentId int) error {
	query := `
	DELETE FROM students
	WHERE id = $1
`
	result, err := r.db.Exec(query, studentId)
	if err != nil {
		return fmt.Errorf("delete student query error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected error: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("student not found")
	}

	return nil
}
