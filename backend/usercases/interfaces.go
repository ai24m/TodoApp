package usercases

import "github.com/ai24m/TodoApp.git/backend/models"

type TodosRepository interface {
	GetAllTodos() ([]models.Todo, error)
}
