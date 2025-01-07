package token

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  int64     `json:"issuedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}



func (p Payload) Valid() error {
	if time.Now().After(p.ExpiresAt) {
		return fmt.Errorf("token has expired")
	}
	return nil
}

func NewPayload(username string, duration time.Duration) (Payload, error) {
	tk, err := uuid.NewRandom()
	return Payload{
		ID:        tk,
		Username:  username,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(duration),
	}, err
}
