package refreshtoken

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() (string, error) {
	// Implement your refresh token generation logic here
	// For example, you can use a random string generator or a UUID generator

	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil

	//return base64.URLEncoding.EncodeToString(b), nil
}
