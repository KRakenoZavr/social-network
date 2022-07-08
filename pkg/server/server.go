package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "social/pkg/middleware"
	"social/pkg/server/ctrl"
	"social/pkg/server/router"
	"strconv"
	"time"
)

type Server struct {
	router *router.Router
}

func NewServer() (s *Server) {
	return &Server{
		router: router.NewRouter(),
	}
}

func (s *Server) Start(port string) error {
	s.configureRouter()

	fmt.Printf("app is running on http://localhost%s\n", port)
	logger := log.New(os.Stdout, "\033[36m", log.LstdFlags)

	c := &ctrl.Controller{Logger: logger, NextRequestID: func() string { return strconv.FormatInt(time.Now().UnixNano(), 36) }}

	server := &http.Server{
		Addr:    port,
		Handler: (ctrl.Middlewares{c.Tracing, c.Logging}).Apply(s.router),
	}

	fmt.Println(s.router.PathMap)

	return server.ListenAndServe()
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", mw.Check(s.load)).Methods("GET")
	s.router.HandleFunc("/:id", mw.Check(s.load)).Methods("GET")
	s.router.HandleFunc("/asd", mw.Check(s.load)).Methods("POST")
}

func (s *Server) load(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/asd" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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
