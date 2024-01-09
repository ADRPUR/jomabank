package util

import (
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min<->max
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner Name
func RandomOwner() string {
	return randomdata.FullName(randomdata.RandomGender)

}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(1, 10000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "MDL", "RON"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
