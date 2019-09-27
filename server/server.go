package server

import (
	"fmt"
	"net/http"
	"os"
)

type Server interface {
	Run() error
}

type server struct {
	addr    string
	handler http.Handler
}

func New(h http.Handler) Server {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	return &server{
		addr:    fmt.Sprintf("%s:%s", host, port),
		handler: h,
	}
}

func (s *server) Run() error {
	return http.ListenAndServe(s.addr, s.handler)
}
