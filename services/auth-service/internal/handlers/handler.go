package handlers

import "github.com/miguelhigueradev/codeflow/services/auth-service/internal/store"

type Handler struct {
	store  *store.Store
	config Config
}

type Config struct {
	GitHubClientID     string
	GitHubClientSecret string
}

func New(s *store.Store, cfg Config) *Handler {
	return &Handler{store: s, config: cfg}
}
