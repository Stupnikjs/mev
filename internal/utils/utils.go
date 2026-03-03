package utils

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func PKFromString(hexKey string) *ecdsa.PrivateKey {
	// Retire le "0x" si présent
	if len(hexKey) >= 2 && hexKey[:2] == "0x" {
		hexKey = hexKey[2:]
	}

	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Fatal("clé privée invalide:", err)
	}

	return privateKey
}
