package server

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"strconv"
)

func createToken(username string, passwordHash string) (string, error) {
	if username == "blockedUser" {
		return "", fmt.Errorf("This user is blocked")
	}
	salt := strconv.Itoa(int(time.Now().UnixNano()))
	secret := username + salt + passwordHash
	hash := sha256.Sum256([]byte(secret))
	token := base64.StdEncoding.EncodeToString(hash[:])
	return token, nil
}
