package modul

type MateriPokokRequest struct {
	Namajabatan  string   `valid:"required"`
	Tugasjabatan []string `valid:"required"`
	Keterampilan []string `valid:"required"`
}
