package namegen

import (
	"crypto/rand"
	"fmt"
	"time"
)

// Generate generates an unique String UUIDv4-like + Current Unix timestamp in HEX.
// It returns the generated String and an any errors.
func Generate() (string, error) {
	var u [16]byte
	now := time.Now()

	// Randomize every byte
	_, err := rand.Read(u[:])
	if err != nil {
		fmt.Println("Couldn't generate a new filename. See error below.")
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:], now.Unix()), nil
}
