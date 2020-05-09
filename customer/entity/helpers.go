package entity

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/google/uuid"
)

// EncryptPassword encrypt string password to sha1 encode
func EncryptPassword(password string) string {
	h := sha1.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

// IsValidUUID to prevent toxic data entering.
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
