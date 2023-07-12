package main

import (
	"net/http"
)

type ApiServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{listenAddr: listenAddr}
}

func (server *ApiServer) Run() {

}

func (server *ApiServer) handleAccount(rw *http.ResponseWriter, req *http.Request) error {
	return nil
}
