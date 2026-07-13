package model

type Subject struct {
	ID            int    `json:"id"`
	SubjectName   string `json:"subject_name"`
	TeacherName   string `json:"teacher_name"`
	TeacherUserID int    `json:"teacher_user_id"`
	Semester      int    `json:"semester"`
}
