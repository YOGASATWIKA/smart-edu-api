package request

type MateriPokokRequest struct {
	Namajabatan  string   `valid:"required"`
	Tugasjabatan []string `valid:"required"`
	Keterampilan []string `valid:"required"`
	Klasifikasi  string   `valid:"required"`
}
