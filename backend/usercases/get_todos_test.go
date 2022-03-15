package usercases_test

import (
	"fmt"
	"testing"

	"github.com/ai24m/TodoApp.git/backend/models"
	"github.com/ai24m/TodoApp.git/backend/usercases"
	"github.com/gomagedon/expectate"
)

//mock data
var codedTodos = []models.Todo{
	{
		Title:       "todo 1",
		Description: "description of todo 1",
		IsComplete:  true,
	},
	{
		Title:       "todo 2",
		Description: "description of todo 2",
		IsComplete:  true,
	},
	{
		Title:       "todo 3",
		Description: "description of todo 3",
		IsComplete:  true,
	},
}

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]models.Todo, error) {
	return nil, fmt.Errorf("List didnt load!") //nolist at all golang has no exceptions, errors are plain values, no try catch
}

type SampleTodosRepo struct{}

func (SampleTodosRepo) GetAllTodos() ([]models.Todo, error) {
	return codedTodos, nil
}

func TestGetTodos(t *testing.T) {
	t.Run("Returns ErrInternal when TodosRepo returns error", func(t *testing.T) {
		//package dependency
		expect := expectate.Expect(t)

		//error cases calling usecase
		repo := new(MockTodosRepo)

		todos, err := usercases.GetTodos(repo)
		// expect(err).ToBe(usercases.ErrInternal)
		expect(todos).ToBe(nil)

		//_, err := usercases.GetTodos(repo) //run test cd into usercases and go test in terminal
		if err != usercases.ErrInternal {
			t.Fatalf("expected ErrInteral; Got: %v", err) //call fail
		}
		if todos != nil {
			t.Fatalf("Expected todos to be nil; Got: %v", todos)
		}
	})

	t.Run("Returns todos from TodoRepo", func(t *testing.T) {
		expect := expectate.Expect(t)
		repo := new(SampleTodosRepo)
		//_, err := usercases.GetTodos(repo)
		todos, err := usercases.GetTodos(repo)
		expect(err).ToBe(nil)
		expect(todos).ToEqual(codedTodos)
	})
}
