package model


type CreateBaseMateriRequest struct {
	Jobid int64 `valid:"required"`
	Namajabatan  string `valid:"required"`
	Tugasjabatan []string `valid:"required"`
	Keterampilan []string `valid:"required"`
	Klasifikasi  string `valid:"required"`
}