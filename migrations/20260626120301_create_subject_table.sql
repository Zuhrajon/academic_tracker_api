-- +goose Up
CREATE TABLE subjects (
    id  BIGSERIAL PRIMARY KEY,
    subject_name VARCHAR(100) NOT NULL,
    teacher_name VARCHAR(100) NOT NULL,
    semester    INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);



-- +goose Down
DROP TABLE subjects;
