package router

import (
	"github.com/ai24m/TodoApp.git/backend/server/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TodoComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTodo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTodo).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTodo).Methods("DELETE", "OPTIONS")
	return router
}
