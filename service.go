package goservice

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Service struct {
	port int
	router *mux.Router
}

func NewService(port int) *Service {
	s := &Service{
		port: port,
		router: mux.NewRouter(),
	}
	s.setRoutes()
	return s
}

func (s *Service) setRoutes() {
	s.router.HandleFunc("/echo/{something}", s.EchoHandler)
}

func (s *Service) EchoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	echo := vars["something"]
	fmt.Fprintf(w, "%s", echo)
}

func (s *Service) Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router))
}
