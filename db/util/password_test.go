package util
import (
	"testing"

	
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	Password := RandomOwner()
	hashedPassword, err := HashPassword(Password)
	if err != nil {
		t.Error(err)
	}
	wrongPass := RandomOwner()
	if err != nil {
		t.Error(err)
	}
	err = CheckPasswordHash(wrongPass, hashedPassword)
	require.Equal(t, err, bcrypt.ErrMismatchedHashAndPassword) 

}