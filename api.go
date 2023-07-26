package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct{ listenAddr string }
type apiFunc func(http.ResponseWriter, *http.Request) error
type ApiError struct{ Error string }

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{listenAddr: listenAddr}
}

func (server *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandler(server.handleAccount))
}

func (server *ApiServer) handleAccount(w *http.ResponseWriter, r *http.Request) error {
	return nil
}

func (server *ApiServer) handleGetAccount(rw *http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *ApiServer) handleCreateAccount(rw *http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *ApiServer) handleUpdateAccount(rw *http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *ApiServer) handleDeleteAccount(rw *http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *ApiServer) handleTransfer(rw *http.ResponseWriter, req *http.Request) error {
	return nil
}
