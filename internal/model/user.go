package model

const (
	RoleAdmin   = "admin"
	RoleTeacher = "teacher"
	RoleStudent = "student"
)

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
	StudentID    int    `json:"student_id"`
}

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	StudentID int    `json:"student_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
