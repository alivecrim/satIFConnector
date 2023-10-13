package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	//s.router.Handle("/api/device", JwtMiddleware(handlers.DeviceEntityHandlerGet(), "USER")).Methods("GET")
	//s.router.Handle("/api/device", JwtMiddleware(handlers.DeviceEntityHandlerPost(), "ADMIN")).Methods("POST")
	//s.router.Handle("/api/device", JwtMiddleware(handlers.DeviceEntityHandlerUpdate(), "ADMIN")).Methods("PUT")
	//s.router.Handle("/api/deviceType", JwtMiddleware(handlers.DeviceTypeEntityHandlerUpdate(), "USER")).Methods("GET")
	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		log.Fatal(err)
	}
}
