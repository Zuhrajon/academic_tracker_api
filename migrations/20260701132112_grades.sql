-- +goose Up
CREATE TABLE IF NOT EXISTS grades (
    id BIGSERIAL PRIMARY KEY,
    student_id BIGINT NOT NULL REFERENCES students(id),
    subject_id BIGINT NOT NULL REFERENCES subjects(id),
    grade_date DATE NOT NULL DEFAULT CURRENT_DATE,
    grade INT NOT NULL CHECK (grade >= 1 AND grade <= 10),
    comment VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- +goose Down
DROP TABLE IF EXISTS grades;