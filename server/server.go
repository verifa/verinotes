package server

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/verifa/verinotes/store"
	"github.com/verifa/verinotes/ui"
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

	if ui.Enabled {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/ui", http.StatusFound)
		})
		r.Mount("/ui", handleUI())
	} else {
		// TODO: this is a bit hacky to work with dev
		// If UI is not enabled, we are likely in dev mode so forward to the
		// default port the frontend runs on in dev mode
		r.Get("/ui", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "http://localhost:5173/ui", http.StatusFound)
		})
	}

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/note", serverImpl.CreateNote)
		r.Get("/notes", serverImpl.QueryAllNotes)
		r.Get("/note/{noteID}", serverImpl.GetNote)
		r.Put("/note/{noteID}", serverImpl.UpdateNote)
		r.Delete("/note/{noteID}", serverImpl.DeleteNote)
	})
	return r, nil
}

// handleUI returns a handler for our Single Page Application that checks if a
// requested resource exists, and if it doesn't, returns the root index.html
// (the single page).
func handleUI() http.Handler {
	index, err := ui.Site.Open("index.html")
	if err != nil {
		log.Fatal("Failed opening UI's index.html: " + err.Error())
	}
	var spaIndex bytes.Buffer
	if _, err := spaIndex.ReadFrom(index); err != nil {
		log.Fatal("Failed reading UI's index.html: " + err.Error())
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Strip the /ui prefix from the requested path to get the path to the
		// requested resource as it would be on the backend filesystem.
		path := strings.TrimPrefix(r.URL.Path, "/ui")
		// If requesting the root page, we will end up with nothing left, so
		// in that case we know it's the root page they were looking for
		if path == "" {
			w.WriteHeader(http.StatusAccepted)
			w.Write(spaIndex.Bytes())
			return
		}
		// Check if requested resource exists. If it does, treat it like a resource
		// such as a .js or .css file with the full path including the filename.
		// If it doesn't exist, it's a path without a filename and we should
		// return our Single Page (index.html)
		f, err := ui.Site.Open(path)
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusAccepted)
			w.Write(spaIndex.Bytes())
			return
		} else if err != nil {
			http.Error(w, "Error: opening requested path "+path+": "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		http.StripPrefix("/ui", http.FileServer(ui.Site)).ServeHTTP(w, r)
	})
}

func returnJSON(w http.ResponseWriter, obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, "Creating JSON response: "+err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Write(b)
}
