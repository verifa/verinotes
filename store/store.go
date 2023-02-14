package store

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/verifa/verinotes/ent"
	"github.com/verifa/verinotes/ent/note"
)

type Store struct {
	ctx context.Context
	//	config *Config
	client *ent.Client
}

type Config struct {
	SessionDuration time.Duration
}

func New(ctx context.Context) (*Store, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	s := Store{
		ctx:    ctx,
		client: client,
	}

	return &s, nil
}

func (s *Store) CreateNote(note *ent.Note) (*ent.Note, error) {
	u, err := s.client.Note.Create().
		SetID(note.ID).
		SetData(note.Data).
		Save(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating note: %w", err)
	}
	log.Println("note was created: ", u)
	return u, nil
}

// TODO never called, should I defer it somewhere (after creating New store?)
func (s *Store) Close() error {
	return s.client.Close()
}

func (s *Store) QueryNote(id int) (*ent.Note, error) {
	u, err := s.client.Note.Query().Where(note.ID(id)).Only(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func (s *Store) QueryAllNotes() ([]*ent.Note, error) {
	u, err := s.client.Note.Query().All(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
