package response

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ModelResponse struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	Model            string             `json:"model" bson:"model"`
	Status           string             `json:"status" bson:"status"`
	PromtContext     string             `json:"promt_context" bson:"promt_context"`
	PromtInstruction string             `json:"promt_instruction" bson:"promt_instruction"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
}
