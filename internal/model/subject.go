package model

type Subject struct {
	ID          int    `json:"id"`
	SubjectName string `json:"subject_name"`
	TeacherName string `json:"teacher_name"`
	Semester    int    `json:"semester"`
}
