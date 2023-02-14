package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/verifa/verinotes/store"
)

type ServerImpl struct {
	store *store.Store
}

func New(ctx context.Context, store *store.Store) (*chi.Mux, error) {
	// Create logger
	logger := httplog.NewLogger("verinotes", httplog.Options{
		JSON:    false, // could expose as config option
		Concise: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/healthz"))

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"*"},
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173", "http://localhost:9998"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		// Debug:            true,
	}))

	serverImpl := ServerImpl{
		store: store,
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/note", serverImpl.CreateNote)
	r.Get("/notes", serverImpl.QueryAllNotes)

	return r, nil
}

func returnJSON(w http.ResponseWriter, obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, "Creating JSON response: "+err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Write(b)
}
