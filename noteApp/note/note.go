package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(noteTitle, noteContent string) (Note, error) {

	if noteTitle == "" || noteContent == "" {
		return Note{}, errors.New(" Title or Content is needed")
	}

	return Note{
		Title:     noteTitle,
		Content:   noteContent,
		CreatedAt: time.Now(),
	}, nil
}

func (anyNote Note) Save() error {

	filename := strings.ReplaceAll(anyNote.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(anyNote)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

// Packges used are strings, time, encoding/json
//  Also add's struct tags for metadata
