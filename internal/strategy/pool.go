package strategy

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

type V2Pool struct {
	Address common.Address
	Active  bool
	Tokens  []string // optionnel : addresses des deux tokens
}

// getReservesABI est le binding minimal pour getReserves()
var getReservesABI = w3.MustNewFunc(
	"getReserves()",
	"(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)",
)

func (p *V2Pool) IsActiveUniswapV2Pair(client *w3.Client) (bool, error) {
	ctx := context.Background()

	// 1. Vérif rapide : est-ce un contrat ?
	var code []byte
	if err := client.CallCtx(ctx, eth.Code(p.Address, nil).Returns(&code)); err != nil {
		return false, fmt.Errorf("getCode failed: %w", err)
	}

	if len(code) == 0 {
		return false, nil // EOA, pas contrat
	}
	// Option 2 : Test vraiment une pair → appel getReserves()
	var (
		reserve0 *big.Int
		reserve1 *big.Int
		ts       uint32
	)
	_ = ts
	err := client.CallCtx(ctx,
		eth.CallFunc(p.Address, getReservesABI).Returns(&reserve0),
	)
	fmt.Println(reserve0)

	if err != nil {
		// Revert ou erreur → probablement pas une pair valide
		log.Printf("getReserves failed on %s: %v", p.Address.Hex(), err)
		return false, nil
	}

	// Optionnel : vérifications supplémentaires pour plus de robustesse
	if reserve0 == nil || reserve1 == nil || reserve0.Sign() < 0 || reserve1.Sign() < 0 {

		return false, nil // réserves invalides
	}

	// Si on arrive ici → c'est très probablement une vraie pair V2
	p.Active = true

	// Bonus : récupérer token0 et token1 si tu veux remplir p.Tokens
	token0ABI := w3.MustNewFunc("token0()", "address")
	token1ABI := w3.MustNewFunc("token1()", "address")

	var token0, token1 common.Address
	_ = client.CallCtx(ctx, // on ignore l'erreur ici car on est déjà confiant
		eth.CallFunc(p.Address, token0ABI).Returns(&token0),
		eth.CallFunc(p.Address, token1ABI).Returns(&token1),
	)

	p.Tokens = []string{token0.Hex(), token1.Hex()}

	return true, nil
}
