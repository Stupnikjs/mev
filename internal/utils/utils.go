package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/Stupnikjs/mev/internal/strategy"
	"github.com/ethereum/go-ethereum/common"
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

func ArgsFromCallData(data []byte) error {
	selector := hex.EncodeToString(data[:4])

	fn := strategy.UniswapV2Funcs[selector]

	switch selector {
	case "0x38ed1739", "0x8803dbee": // tokens→tokens
		var amountA, amountB *big.Int
		var path []common.Address
		var to common.Address
		var deadline *big.Int

		err := fn.DecodeArgs(data, &amountA, &amountB, &path, &to, &deadline)
		if err == nil {
			fmt.Println(amountA, amountB, path, to, deadline)
		}
	case "0x7ff36ab5": // ETH→tokens
		var amountOutMin *big.Int
		var path []common.Address
		var to common.Address
		var deadline *big.Int

		_ = fn.DecodeArgs(data, &amountOutMin, &path, &to, &deadline)
		fmt.Println(amountOutMin, path, to, deadline)
	// note: amountIn = tx.Value()
	default:
		return nil
	}
	return nil
}
