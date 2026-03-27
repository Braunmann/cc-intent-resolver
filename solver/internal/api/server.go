package api

import (
	"encoding/json"
	"net/http"
	"solver/internal/store"
)

type Server struct {
	store *store.IntentStore
	port  string
}

func NewServer(store *store.IntentStore, port string) (*Server, error) {
	return &Server{
		store: store,
		port:  port,
	}, nil
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/intents", s.handleGetIntents)

	return http.ListenAndServe(s.port, mux)
}

func (s *Server) handleGetIntents(w http.ResponseWriter, r *http.Request) {
	intents, err := s.store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(intents)
}
