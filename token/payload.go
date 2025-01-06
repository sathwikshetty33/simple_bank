package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  int64     `json:"issuedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func NewPayload(username string, duration time.Duration) (Payload, error) {
	tk,err := uuid.NewRandom()
	return Payload{
		ID:        tk,
		Username:  username,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(duration),
	}, err
}