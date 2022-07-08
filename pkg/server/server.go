package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social/pkg/server/router"
)

type Server struct {
	router *router.Router
}

func NewServer() (s *Server) {
	s = &Server{
		router: router.NewRouter(),
	}

	return s
}

func (s *Server) Start(bindAddr string) error {
	s.configureRouter()

	fmt.Printf("app is running on %s\n", bindAddr)

	return http.ListenAndServe(bindAddr, s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.load).Methods("GET")
}

func (s *Server) load(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Lol bool
		Kek string
	}{
		true,
		"user",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
