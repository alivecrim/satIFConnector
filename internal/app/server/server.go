package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
	Logger *logrus.Logger
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
		Logger: logrus.New(),
	}
}

func (s *Server) Init() {
	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		log.Fatal(err)
	}
}
