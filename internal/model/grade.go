package model

type Grade struct {
	ID        int    `json:"id"`
	StudentId int    `json:"student_id"`
	SubjectId int    `json:"subject_id"`
	GradeDate string `json:"grade_date"`
	Grade     int    `json:"grade"`
	Comment   string `json:"comment"`
}
