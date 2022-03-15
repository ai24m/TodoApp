package ui_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ai24m/TodoApp.git/backend/models"
	"github.com/ai24m/TodoApp.git/backend/ui"
	"github.com/gomagedon/expectate"
)

// MockService
type MockService struct {
	err   error
	todos []models.Todo
}

func (mockSvc MockService) GetAllTodos() ([]models.Todo, error) {
	if mockSvc.err != nil {
		return nil, mockSvc.err
	}
	return mockSvc.todos, nil
}

// Hard codecd Todos
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

//HTTPTests
type HTTPTest struct {
	name string

	service     *MockService
	inputMethod string
	inputURL    string

	expectedErr    error
	expectedStatus int
	expectedTodos  []models.Todo
}

//Main Test
func TestHTTP(t *testing.T) {
	tests := getTests()
	tests = append(tests, getDisallowedMethodTests()...)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testHTTP(t, test)
		})
	}
}

func testHTTP(t *testing.T, test HTTPTest) {
	expect := expectate.Expect(t)

	// service := &MockService{
	// 	err: fmt.Errorf("something internal went wrong"),
	// }
	w := httptest.NewRecorder()
	//r := httptest.NewRequest(http.MethodGet, "http://mytodoapp.com/", nil)
	r := httptest.NewRequest(test.inputMethod, test.inputURL, nil)

	server := ui.NewHTTP()
	server.UseService(test.service)
	//server.UserService(service)
	server.ServeHTTP(w, r)

	var body []models.Todo
	json.NewDecoder(w.Result().Body).Decode(&body)

	expect(w.Result().StatusCode).ToBe(test.expectedStatus)
	expect(body).ToEqual(test.expectedTodos)
	//expect(w.Result().StatusCode).ToBe(http.StatusInternalServerError)

}

func getTests() []HTTPTest {
	return []HTTPTest{
		{
			name:           "Gives 500 and no todos",
			service:        &MockService{err: fmt.Errorf("internal status ")},
			inputMethod:    "GET",
			inputURL:       "http://mytodoapp.com/todos",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Gives 404 wrong path",
			service:        &MockService{todos: codedTodos},
			inputMethod:    "GET",
			inputURL:       "http://mytodoapp.com/foo",
			expectedStatus: 404,
			expectedTodos:  nil,
		},
		{
			name:           "Gives 405 wrong method",
			service:        &MockService{todos: codedTodos},
			inputMethod:    "GET",
			inputURL:       "http://mytodoapp.com/bar",
			expectedStatus: 404,
			expectedTodos:  nil,
		},
		{
			name:           "Service returns todos if no error",
			service:        &MockService{todos: codedTodos},
			inputMethod:    "GET",
			inputURL:       "http://mytodoapp.com/todos",
			expectedStatus: 200,
			expectedTodos:  codedTodos,
		},
	}

}

func getDisallowedMethodTests() []HTTPTest {
	tests := []HTTPTest{}

	disallowedmethods := []string{
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
	}

	for _, method := range disallowedmethods {
		tests = append(tests, HTTPTest{
			name:           fmt.Sprintf("Method %s gives 405 status, no todos", method),
			service:        &MockService{todos: codedTodos},
			inputURL:       "http://mytodoapp.com/todos/",
			inputMethod:    method,
			expectedStatus: http.StatusMethodNotAllowed,
		})
	}
	return tests
}
