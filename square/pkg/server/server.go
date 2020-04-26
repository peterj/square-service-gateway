package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server wraps the Mux router
type Server struct {
	Mux *mux.Router
}

// New creates a new instance of the Server
func New(ctx context.Context) *Server {
	s := &Server{
		Mux: mux.NewRouter(),
	}
	s.Mux.Use(WithLogging)

	s.Mux.Handle("/metrics", promhttp.Handler())
	s.Mux.HandleFunc("/square/{number}", squareHandler).Methods("GET")
	return s
}

// ServeHTTP is an HTTP handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}
