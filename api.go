package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{listenAddr: listenAddr}
}

func (server *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc()
}

func (server *ApiServer) handleAccount(rw *http.ResponseWriter, req *http.Request) error {
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
