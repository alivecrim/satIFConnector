package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pasha.com/satifconnector/internal/app/server/handlers/glob"
	groundstation_http "pasha.com/satifconnector/internal/app/server/handlers/groundstation"
	project_http "pasha.com/satifconnector/internal/app/server/handlers/project"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

func (s *Server) Init() {
	s.router.Handle("/api/relationState", glob.RequestForCurrentRelationState()).Methods("GET")
	s.router.Handle("/api/crud/project", project_http.Create()).Methods("POST")
	s.router.Handle("/api/crud/project/setActive/{id}", project_http.SetActive()).Methods("PUT")

	s.router.Handle("/api/crud/groundStation", groundstation_http.Create()).Methods("POST")
	s.router.Handle("/api/crud/groundStation", groundstation_http.FindAll()).Methods("GET")
	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		log.Fatal(err)
	}
}
