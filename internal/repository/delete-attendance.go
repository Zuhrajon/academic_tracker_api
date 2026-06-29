package repository

import "fmt"

func (r *Repository) DeleteAttendance(attendanceId int) error {
	query := `
	DELETE FROM attendance
	WHERE id = $1
`
	result, err := r.db.Exec(query, attendanceId)
	if err != nil {
		return fmt.Errorf("delete attendance query error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected error: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("attendance not found")
	}

	return nil
}
