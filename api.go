package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiError struct{ Error string }
type ApiServer struct{ listenAddr string }
type apiFunc func(http.ResponseWriter, *http.Request) error

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
	router.HandleFunc("/account/{id}", makeHTTPHandler(server.handleGetAccount))
	router.HandleFunc("/account/transfer", makeHTTPHandler(server.handleTransfer))

	log.Println("Server running on port", server.listenAddr)

	http.ListenAndServe(server.listenAddr, router)
}

func (server *ApiServer) handleAccount(rw http.ResponseWriter, req *http.Request) error {
	if req.Method == "GET" {
		return server.handleListAccounts(rw, req)
	}

	if req.Method == "POST" {
		return server.handleCreateAccount(rw, req)
	}

	if req.Method == "PUT" {
		return server.handleUpdateAccount(rw, req)
	}

	if req.Method == "DELETE" {
		return server.handleDeleteAccount(rw, req)
	}

	return fmt.Errorf("method not allowed %s", req.Method)
}

func (server *ApiServer) handleListAccounts(rw http.ResponseWriter, req *http.Request) error {
	accounts := &Account{}

	return WriteJSON(rw, http.StatusOK, accounts)
}

func (server *ApiServer) handleGetAccount(rw http.ResponseWriter, req *http.Request) error {
	account := mux.Vars(req)["id"]

	return WriteJSON(rw, http.StatusOK, account)
}

func (server *ApiServer) handleCreateAccount(rw http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *ApiServer) handleUpdateAccount(rw http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *ApiServer) handleDeleteAccount(rw http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *ApiServer) handleTransfer(rw http.ResponseWriter, req *http.Request) error {
	return nil
}
