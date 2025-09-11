package dto

type CreateTodoDTO struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoDTO struct {
    Title     string `json:"title" binding:"required"`
    Completed bool   `json:"completed"`
}