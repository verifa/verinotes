package server

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"github.com/verifa/verinotes/store"
)

var (
	defaultNote = "test note content"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	// s.Router.ServeHTTP(rr, req)
	s.ServeHTTP(rr, req)
	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func serverSetup(t *testing.T) *chi.Mux {
	ctx := context.TODO()
	// Create a New Server Struct

	dbFileName := t.Name()

	store, err := store.NewTest(ctx, dbFileName)
	if err != nil {
		t.Errorf("creating store: %v", err)
	}

	srv, err := New(ctx, store)
	if err != nil {
		t.Errorf("creating server: %v", err)
	}

	return srv
}

func createOneNote(t *testing.T, srv *chi.Mux, note string) string {
	rawBody := `{"data":"` + note + `"}`
	// Create a New Request
	req, _ := http.NewRequest("POST", "/api/v1/note", bytes.NewBuffer([]byte(rawBody)))

	// Execute Request
	response := executeRequest(req, srv)

	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)

	// We can use testify/require to assert values, as it is more convenient
	return response.Body.String()
}

func TestCreateOneNote(t *testing.T) {
	srv := serverSetup(t)

	resp := createOneNote(t, srv, defaultNote)

	expectedResp := `{"id":1,"data":"test note content"}`
	// We can use testify/require to assert values, as it is more convenient
	require.Equal(t, expectedResp, resp)

}

func TestQueryAllNotes(t *testing.T) {
	// Create one note that we can query
	TestCreateOneNote(t)

	srv := serverSetup(t)
	// Create a New Request
	req, _ := http.NewRequest("GET", "/api/v1/notes", nil)

	// Execute Request
	response := executeRequest(req, srv)

	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)

	expectedResp := `[{"id":1,"data":"test note content"}]`
	// We can use testify/require to assert values, as it is more convenient
	require.Equal(t, expectedResp, response.Body.String())
}

func TestRemoveNote(t *testing.T) {
	// Create one note that we can query
	srv := serverSetup(t)

	createOneNote(t, srv, defaultNote)
	createOneNote(t, srv, defaultNote)

	// Create a New Request
	req, _ := http.NewRequest("DELETE", "/api/v1/note/1", nil)

	// Execute Request
	response := executeRequest(req, srv)

	// Check the response code, no body will be returned from deletes, only the status 200
	checkResponseCode(t, http.StatusOK, response.Code)

	// now test we still have the note with id 2

	req, _ = http.NewRequest("GET", "/api/v1/notes", nil)
	response = executeRequest(req, srv)
	checkResponseCode(t, http.StatusOK, response.Code)

	expectedResp := `[{"id":2,"data":"test note content"}]`
	require.Equal(t, expectedResp, response.Body.String())
}

func TestUpdateNote(t *testing.T) {
	// Create one note that we can query
	srv := serverSetup(t)

	createOneNote(t, srv, defaultNote)

	note := "modified note content"
	rawBody := `{"data":"` + note + `"}`
	// Create a New Request
	req, _ := http.NewRequest("PUT", "/api/v1/note/1", bytes.NewBuffer([]byte(rawBody)))

	// Execute Request
	response := executeRequest(req, srv)

	// Check the response code, no body will be returned from deletes, only the status 200
	checkResponseCode(t, http.StatusOK, response.Code)

	// now test we still have the note with id 2

	req, _ = http.NewRequest("GET", "/api/v1/note/1", nil)
	response = executeRequest(req, srv)
	checkResponseCode(t, http.StatusOK, response.Code)

	expectedResp := `{"id":1,"data":"` + note + `"}`
	require.Equal(t, expectedResp, response.Body.String())
}
