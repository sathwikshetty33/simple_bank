package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  int64     `json:"issuedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}



func (p *Payload) Valid() error {
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
