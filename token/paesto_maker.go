package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)
type PaestoMaker struct{
	paesto *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error){
	if len(symmetricKey) != chacha20poly1305.KeySize{
		return nil, fmt.Errorf("symmetric key must be 32 bytes long")
	}
	maker := &PaestoMaker{
		paesto: paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		}
		return maker, nil
}

func (maker *PaestoMaker) CreateToken(username string, duration time.Duration) (string,*Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "",&payload,err
	}
	token, err := maker.paesto.Encrypt( maker.symmetricKey,payload ,nil)
	if err != nil {
		return "",&payload,err
	}
	return token,&payload, nil
}



func(maker *PaestoMaker) VerifyToken(token string) (*Payload, error){
	payload := &Payload {}
	err := maker.paesto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}



