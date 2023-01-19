package model

import (
	"github.com/google/uuid"
)

type Mensagem struct {
	ID   uuid.UUID `json:"Id" db:"_id" validate:"omitempty"`
	Body string    `json:"body" db:"body" validate:"required"`
}
