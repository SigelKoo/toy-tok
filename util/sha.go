package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetKeccakStr(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
