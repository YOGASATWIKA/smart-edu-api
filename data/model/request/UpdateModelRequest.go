package request

type UpdateModelRequest struct {
	Nama             string `valid:"required"`
	Model            string `valid:"required"`
	PromtContext     string `valid:"required"`
	PromtInstruction string `valid:"required"`
}
