package repository

import "fmt"

func (r *Repository) DeleteGrade(gradeID int) error {
	query := `
	DELETE FROM grades
	WHERE id = $1
`
	result, err := r.db.Exec(query, gradeID)
	if err != nil {
		return fmt.Errorf("delete grade query error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected error: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("grade not found")
	}

	return nil
}
