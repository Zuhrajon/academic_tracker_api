package repository

import "academic-tracker-api/internal/model"

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetStudents() []model.Student {
	return []model.Student{
		{ID: 1, FirstName: "Zuzu", LastName: "...", GroupName: "GO-21", Course: 1},
		{ID: 2, FirstName: "hihih", LastName: "...", GroupName: "GO-23", Course: 1},
	}
}
