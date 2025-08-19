package request


type CreateBaseMateriRequest struct {
	Namajabatan  string `valid:"required"`
	Tugasjabatan []string `valid:"required"`
	Keterampilan []string `valid:"required"`
	Klasifikasi  string `valid:"required"`
	Status  string `valid:"required"`
}