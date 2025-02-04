package util

import (
	"math/rand"
	"strings"
	"time"
)
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func RandomString(n int) string {
	var sb strings.Builder 
	k := len(alphabet)
	for i:=0;i<n;i++{
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}
func RandomMoney() int64 {
	return int64(RandomInt(0, 100000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}