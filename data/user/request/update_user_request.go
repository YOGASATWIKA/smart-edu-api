package request

type UpdateUserRequest struct {
	Email   string `json:"email" validate:"required"`
	Picture string `json:"picture" validate:"required"`
	Name    string `json:"name" validate:"required"`
}
