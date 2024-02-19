package entity

import "github.com/google/uuid"

type Profile struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}
