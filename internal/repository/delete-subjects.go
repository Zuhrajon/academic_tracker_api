package repository

import "fmt"

func (r *Repository) DeleteSubjects(subjectId int) error {
	query := `
	DELETE FROM subjects
	WHERE id = $1
`
	result, err := r.db.Exec(query, subjectId)
	if err != nil {
		return fmt.Errorf("delete subject query error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected error: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("subject not found")
	}

	return nil
}
