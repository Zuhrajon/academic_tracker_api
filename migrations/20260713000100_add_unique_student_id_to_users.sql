-- +goose Up
UPDATE users
SET student_id = NULL
WHERE student_id = 0;

CREATE UNIQUE INDEX IF NOT EXISTS users_student_id_unique_idx
ON users(student_id)
WHERE student_id IS NOT NULL;

-- +goose Down
DROP INDEX IF EXISTS users_student_id_unique_idx;
