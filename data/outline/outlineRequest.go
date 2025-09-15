package outline

import (
	"smart-edu-api/embeded"
	"smart-edu-api/entity"
)

type OutlineRequest struct {
	MateriPokok *entity.MateriPokok
	Outline     *embeded.Outline
	Err         error
}
type ModelRequest struct {
	Model string `json:"model"`
}
