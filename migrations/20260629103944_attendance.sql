-- +goose Up
CREATE TABLE IF NOT EXISTS attendance (
    id BIGSERIAL PRIMARY KEY,
    student_id BIGINT NOT NULL REFERENCES students(id),
    subject_id BIGINT NOT NULL REFERENCES subjects(id),
    lesson_date DATE NOT NULL DEFAULT CURRENT_DATE,
    status VARCHAR(100) CHECK (status IN ('present', 'absent', 'late')) NOT NULL,
    comment VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS attendance;
