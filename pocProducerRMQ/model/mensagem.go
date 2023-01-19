package model

import (
	"github.com/google/uuid"
)

type Mensagem struct {
	ID   uuid.UUID `json:"Id" db:"_id" validate:"omitempty"`
	Body string    `json:"body" db:"body"`
}

type MensagemErro struct {
	ID    uuid.UUID `json:"Id" db:"_id" validate:"omitempty"`
	Corpo string    `json:"corpo" db:"corpo"`
	Idade int       `json:"idade" db:"idade"`
}
