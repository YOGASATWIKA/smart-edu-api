package modul

type ModelRequest struct {
	Model string   `valid:"required"`
	Id    []string `valid:"required"`
}
