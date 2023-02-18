package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/verifa/verinotes/ent"
)

func (s *ServerImpl) CreateNote(w http.ResponseWriter, r *http.Request) {
	req := ent.Note{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	note, err := s.store.CreateNote(&req)
	if err != nil {
		http.Error(w, "Creating note: "+err.Error(), http.StatusBadRequest)
		return
	}
	returnJSON(w, note)
}

func (s *ServerImpl) UpdateNote(w http.ResponseWriter, r *http.Request) {
	req := ent.Note{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	noteParam, err := strconv.Atoi(chi.URLParam(r, "noteID"))
	if err != nil {
		http.Error(w, "Converting noteID to int: "+err.Error(), http.StatusBadRequest)
		return
	}
	note, err := s.store.UpdateNote(noteParam, &req)
	if err != nil {
		http.Error(w, "Updating note: "+err.Error(), http.StatusBadRequest)
		return
	}
	returnJSON(w, note)
}

func (s *ServerImpl) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteParam, err := strconv.Atoi(chi.URLParam(r, "noteID"))
	if err != nil {
		http.Error(w, "Converting noteID to int: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = s.store.DeleteNote(noteParam)
	if err != nil {
		http.Error(w, "Deleting note: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *ServerImpl) QueryAllNotes(w http.ResponseWriter, r *http.Request) {
	note, err := s.store.QueryAllNotes()
	if err != nil {
		http.Error(w, "Creating note: "+err.Error(), http.StatusBadRequest)
		return
	}
	returnJSON(w, note)
}

func (s *ServerImpl) GetNote(w http.ResponseWriter, r *http.Request) {
	noteParam, err := strconv.Atoi(chi.URLParam(r, "noteID"))
	if err != nil {
		http.Error(w, "Converting noteID to int: "+err.Error(), http.StatusBadRequest)
		return
	}
	note, err := s.store.QueryNote(noteParam)
	if err != nil {
		http.Error(w, "Getting note: "+err.Error(), http.StatusBadRequest)
		return
	}
	returnJSON(w, note)
}

func (s *ServerImpl) CloseDB() {
	s.store.Close()
}
