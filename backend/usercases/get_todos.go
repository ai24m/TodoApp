package usercases

import "github.com/ai24m/TodoApp.git/backend/models"

func GetTodos(repo TodosRepository) ([]models.Todo, error) {
	//_, err := repo.GetAllTodos() //just check if error passes
	todos, err := repo.GetAllTodos()
	if err != nil {
		return nil, ErrInternal
	}
	//return nil, nil
	return todos, nil
}
