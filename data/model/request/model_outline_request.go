package request

import "smart-edu-api/entity"

type ModelOutlineRequest struct {
	Model       string              `json:"model" validate:"required"`
	Description string              `json:"description" validate:"required"`
	Steps       []entity.PromptStep `json:"steps" validate:"required,dive"`
	Variables   []string            `json:"variables" validate:"required,dive"`
	IsActive    bool                `json:"is_active"`
}
