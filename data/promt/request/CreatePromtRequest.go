package request

type CreatePromtRequest struct {
	Nama string `valid:"required"`
	Model  string `valid:"required"`
	Status string `valid:"required"`
	PromtContext string `valid:"required"`
	PromtInstruction string `valid:"required"`
}