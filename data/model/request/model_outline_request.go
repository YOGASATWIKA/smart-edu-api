package request

type ModelOutlineRequest struct {
	Model string `json:"model" valid:"required"`
	Promt Promt  `json:"promt" valid:"required"`
}
type Promt struct {
	SystemPrompt string   `json:"system_prompt" valid:"required"`
	UserPrompts  []string `json:"user_prompts" valid:"required"`
}
