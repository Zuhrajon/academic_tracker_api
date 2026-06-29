package model

type Attendance struct {
	ID         int    `json:"id"`
	StudentId  int    `json:"student_id"`
	SubjectId  int    `json:"subject_id"`
	LessonDate string `json:"lesson_date"`
	Status     string `json:"status"`
	Comment    string `json:"comment"`
}
