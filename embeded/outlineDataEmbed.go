package embeded

type Outline struct {
	ListMateri []MateriPokok `json:"list_materi" bson:"list_materi"`
}

type MateriPokok struct {
	MateriPokok   string           `json:"materi_pokok" bson:"materi_pokok"`
	ListSubMateri []SubMateriPokok `json:"list_sub_materi" bson:"list_sub_materi"`
}

type SubMateriPokok struct {
	SubMateriPokok string   `json:"sub_materi_pokok" bson:"sub_materi_pokok"`
	ListMateri     []string `json:"list_materi" bson:"list_materi"`
}
