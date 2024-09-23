package functions

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func ValidateSignature(payload []byte, signature, webhookSecret string) bool {
	mac := hmac.New(sha1.New, []byte(webhookSecret))
	mac.Write(payload)
	expectedSignature := "sha1=" + hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}
