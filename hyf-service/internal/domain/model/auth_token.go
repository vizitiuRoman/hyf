package model

import "github.com/gofrs/uuid"

type AuthToken struct {
	UUID   uuid.UUID
	Token  string
	UserID int64

	// ExpiresIn --> millisecond
	ExpiresIn int64
}
