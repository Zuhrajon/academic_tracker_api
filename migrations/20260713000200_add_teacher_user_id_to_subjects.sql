-- +goose Up
ALTER TABLE subjects
ADD COLUMN IF NOT EXISTS teacher_user_id BIGINT REFERENCES users(id);

CREATE INDEX IF NOT EXISTS subjects_teacher_user_id_idx
ON subjects(teacher_user_id);

-- +goose Down
DROP INDEX IF EXISTS subjects_teacher_user_id_idx;

ALTER TABLE subjects
DROP COLUMN IF EXISTS teacher_user_id;
