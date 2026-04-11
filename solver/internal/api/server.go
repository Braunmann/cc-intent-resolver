package api

import (
	"encoding/json"
	"net/http"
	"solver/internal/config"
	"solver/internal/price"
	"solver/internal/store"

	"github.com/ethereum/go-ethereum/common"
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

	mux.HandleFunc("/v1/intents", s.handleGetIntents)
	mux.HandleFunc("/v1/intents/{id}", s.handleGetIntent)
	mux.HandleFunc("/v1/prices", s.handleGetPrices)
	mux.HandleFunc("/v1/config", s.handleGetConfig)

	handler := corsMiddleware(mux)

	return http.ListenAndServe(s.port, handler)
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

func (s *Server) handleGetIntent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	hash := common.HexToHash(id)

	intent, ok := s.store.Get(hash)
	if !ok {
		http.Error(w, "Intent not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(intent)
}

func (s *Server) handleGetPrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(price.GetPrices())
}

func (s *Server) handleGetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(config.App)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
