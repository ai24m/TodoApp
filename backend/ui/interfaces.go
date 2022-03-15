package ui

import "github.com/ai24m/TodoApp.git/backend/models"

type Service interface {
	GetAllTodos() ([]models.Todo, error)
}
