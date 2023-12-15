package pionex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func generateSignature(key, data string) string {
    keyBytes := []byte(key)
	dataBytes := []byte(data)

	// Create an HMAC-SHA256 hasher
	hasher := hmac.New(sha256.New, keyBytes)

	// Write the data to the hasher
	hasher.Write(dataBytes)

	// Get the HMAC-SHA256 hash
	hash := hasher.Sum(nil)

	// Encode the hash to a hex string
	hashString := hex.EncodeToString(hash)

	return hashString
}
