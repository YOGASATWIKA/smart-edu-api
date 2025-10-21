package entity

import (
	"encoding/json"
	"io"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ebook struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Parts       []*Part            `json:"parts" bson:"parts"`
	Lock        *sync.Mutex        `json:"-" bson:"lock"`
	HtmlContent string             `json:"html_content,omitempty" bson:"html_content,omitempty"`
	JsonContent interface{}        `json:"json_content,omitempty" bson:"json_content,omitempty"`
	ModuleId    primitive.ObjectID `json:"modul" bson:"modul"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt    time.Time          `json:"delete_at" bson:"delete_at,omitempty"`
}

type Part struct {
	Subject       string     `json:"subject" bson:"subject"`
	Introductions []string   `json:"introductions" bson:"introductions"`
	Urgencies     []string   `json:"urgencies" bson:"urgencies"`
	Chapters      []*Chapter `json:"chapters" bson:"chapters"`
}

type Chapter struct {
	Title            string      `json:"title" bson:"title"`
	BaseCompetitions []string    `json:"base_competitions" bson:"base_competitions"`
	TriggerQuestions []string    `json:"trigger_questions" bson:"trigger_questions"`
	Materials        []*Material `json:"materials" bson:"materials"`
	Conclusion       string      `json:"conclusion" bson:"conclusion"`
	Reflections      []string    `json:"reflections" bson:"reflections"`
}

type Material struct {
	Title   string    `json:"title" bson:"title"`
	Short   string    `json:"short" bson:"short"`
	Details []*Detail `json:"details" bson:"details"`
}

type Detail struct {
	Content      string   `json:"content" bson:"content"'`
	Expanded     string   `json:"expanded" bson:"expanded"`
	ExpandChunks []string `json:"expand_chunks" bson:"expand_chunks"`
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
