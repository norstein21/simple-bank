package util

import (
	"math/rand"

	"time"
)
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init(){
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt gemerates a random number between min and max
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Int63n(int64(len(alphabet)))]
	}
	return string(b)
	
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomBalannce() int64{
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string{
	// List of currencies in the world
	currencies := []string{"EUR", "USD", "JPY","KRW"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
	
}