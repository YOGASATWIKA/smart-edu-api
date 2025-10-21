package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Modul struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	MateriPokok MateriPokok        `json:"materi_pokok" bson:"materi_pokok"`
	Outline     Outline            `json:"outline,omitempty" bson:"outline,omitempty"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	State       string             `json:"state" bson:"state"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt    time.Time          `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
}

// Materi Pokok Embeded
type MateriPokok struct {
	Namajabatan  string   `json:"nama_jabatan" bson:"nama_jabatan"`
	Tugasjabatan []string `json:"tugas_jabatan" bson:"tugas_jabatan"`
	Keterampilan []string `json:"keterampilan" bson:"keterampilan"`
}

// Outline Embeded
type Outline struct {
	ListMateri []Materi `json:"list_materi" bson:"list_materi"`
}

type Materi struct {
	MateriPokok   string           `json:"materi_pokok" bson:"materi_pokok"`
	ListSubMateri []SubMateriPokok `json:"list_sub_materi" bson:"list_sub_materi"`
}

type SubMateriPokok struct {
	SubMateriPokok string   `json:"sub_materi_pokok" bson:"sub_materi_pokok"`
	ListMateri     []string `json:"list_materi" bson:"list_materi"`
}
