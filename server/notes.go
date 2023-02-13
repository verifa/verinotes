package server

import (
	"encoding/json"
	"net/http"

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
