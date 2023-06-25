package http

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/go-chi/chi"
)

type Handler struct {
	Router *chi.Mux
	DB     *badger.DB
}

type Request struct {
	Body string `json:"body"`
}

func NewHandler() *Handler {
	return &Handler{
		Router: chi.NewMux(),
	}
}

func (h *Handler) SetupRoutes() {
	router := h.Router
	router.Get("/ping", h.Ping)
	router.Get("/{uri}", h.Redirect)
	router.Post("/shorten", h.ShortenURI)
}
