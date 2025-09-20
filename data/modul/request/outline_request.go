package modul

import (
	"smart-edu-api/entity"
)

type OutlineRequest struct {
	MateriPokok *entity.MateriPokok
	Outline     *entity.Outline
	Err         error
}
