package modul

import (
	"smart-edu-api/embeded"
)

type OutlineRequest struct {
	MateriPokok *embeded.MateriPokok
	Outline     *embeded.Outline
	Err         error
}
type ModelRequest struct {
	Model string `json:"model"`
}
