package entity

import (
	"encoding/json"
	"io"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ebook struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Parts     []*Part            `json:"parts"`
	Lock      *sync.Mutex        `json:"-"`
	Type      string             `json:"type" bson:"type"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt  time.Time          `json:"delete_at" bson:"delete_at,omitempty"`
}

// Comment: Bab
type Part struct {
	Subject       string     `json:"subject"`
	Introductions []string   `json:"introductions"`
	Urgencies     []string   `json:"urgencies"`
	Chapters      []*Chapter `json:"chapters"`
}

type Chapter struct {
	Title            string      `json:"title"`
	BaseCompetitions []string    `json:"base_competitions"`
	TriggerQuestions []string    `json:"trigger_questions"`
	Materials        []*Material `json:"materials"`
	Conclusion       string      `json:"conclusion"`
	Reflections      []string    `json:"reflections"`
}

type Material struct {
	Title   string    `json:"title"`
	Short   string    `json:"short"`
	Details []*Detail `json:"details"`
}

type Detail struct {
	Content      string   `json:"content"`
	Expanded     string   `json:"expanded"`
	ExpandChunks []string `json:"expand_chunks"`
}

func (e *Ebook) Save(file io.Writer) error {
	dt, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}

	e.Lock.Lock()

	defer e.Lock.Unlock()
	_, err = file.Write(dt)
	return err
}
