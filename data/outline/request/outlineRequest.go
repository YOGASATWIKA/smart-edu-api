package outline

type Outline struct {
	ListMateri  []Materi `json:"list_materi" bson:"list_materi"`
}
type Materi struct {
	MateriPokok   string      `json:"materi_pokok" bson:"materi_pokok"`
	ListSubMateri []SubMateri `json:"list_sub_materi" bson:"list_sub_materi"`
}
type SubMateri struct {
	SubMateriPokok string   `json:"sub_materi_pokok" bson:"sub_materi_pokok"`
	ListMateri     []string `json:"list_materi" bson:"list_materi"`
}
