package utils

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/mr-tron/base58"
)

func VerifyMessage(message string, signedMessage string, pubkey string) bool {
	bytes, err := base58.Decode(pubkey)
	if err != nil {
		return false
	}

	messageAsBytes := []byte(message)

	signedMessageAsBytes, err := hex.DecodeString(signedMessage)
	if err != nil {
		return false
	}

	return ed25519.Verify(bytes, messageAsBytes, signedMessageAsBytes)
}
