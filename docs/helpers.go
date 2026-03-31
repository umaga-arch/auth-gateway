package authgateway

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

// GenerateToken generates a cryptographically secure random token.
func GenerateToken(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err!= nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// GenerateTokenWithExpiration generates a cryptographically secure random token with a specified expiration time.
func GenerateTokenWithExpiration(length int, expiration time.Time) (string, error) {
	token, err := GenerateToken(length)
	if err!= nil {
		return "", err
	}
	return token + "." + expiration.UTC().Format(time.RFC3339), nil
}