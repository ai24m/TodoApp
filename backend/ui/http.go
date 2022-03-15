package ui

import (
	"encoding/json"
	"net/http"
)

type HTTPServer struct {
	svc Service
}

func NewHTTP() *HTTPServer { //* operator to hold memory address and resolve it, otherwise type declaration
	return &HTTPServer{} //& operator to get memory address
}

func (server *HTTPServer) UseService(serv Service) {
	server.svc = serv
}

func (server HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	verified := server.verifyRequest(w, r)
	if !verified {
		return
	}

	todos, err := server.svc.GetAllTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(todos)

	w.WriteHeader(http.StatusInternalServerError) //returns 500 instead of default 200
}

func (HTTPServer) verifyRequest(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/todos/" && r.URL.Path != "/todos" {
		w.WriteHeader(http.StatusNotFound)
		return false
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}
