package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/Stupnikjs/mev/config"
	"github.com/Stupnikjs/mev/internal/client"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

var (
	// Fonction ABI déclarée directement en Go, pas de fichier JSON
	funcGetReserves = w3.MustNewFunc(
		"getReserves()",
		"uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast",
	)

	// Pool WETH/USDC Uniswap V2 mainnet
	// token0 = USDC, token1 = WETH
	addrPair = w3.A("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")
	rpc_pub  = "https://ethereum-rpc.publicnode.com"
	pk       = os.Getenv("PRIVATE_KEY")
)

func main() {
	clients, err := client.New(rpc_pub, config.RelayMainnet, pk)
}

func example() {
	// 1. Connexion RPC — RPC public gratuit, pas besoin d'Alchemy pour ce test
	client, err := w3.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		log.Fatal("connexion RPC échouée:", err)
	}
	defer client.Close()

	// 2. Variables qui vont recevoir les résultats
	var (
		reserve0           big.Int // USDC (6 décimales)
		reserve1           big.Int // WETH (18 décimales)
		blockTimestampLast uint32
	)

	// 3. eth_call → lit les réserves du pool
	// w3 batch la requête automatiquement (1 seul round-trip)
	if err := client.Call(
		eth.CallFunc(addrPair, funcGetReserves).
			Returns(&reserve0, &reserve1, &blockTimestampLast),
	); err != nil {
		log.Fatal("appel getReserves échoué:", err)
	}

	// 4. Affiche les résultats
	// USDC a 6 décimales → divise par 1e6
	// WETH a 18 décimales → utilise w3.FromWei
	usdcReserve := new(big.Float).Quo(
		new(big.Float).SetInt(&reserve0),
		big.NewFloat(1e6),
	)

	fmt.Printf("Pool WETH/USDC Uniswap V2\n")
	fmt.Printf("  Adresse  : %s\n", addrPair)
	fmt.Printf("  USDC     : %.2f USDC\n", usdcReserve)
	fmt.Printf("  WETH     : %s WETH\n", w3.FromWei(&reserve1, 18))
	fmt.Printf("  Prix implicite : 1 WETH = %.2f USDC\n",
		new(big.Float).Quo(usdcReserve, new(big.Float).SetFloat64(
			func() float64 {
				f, _ := new(big.Float).SetInt(&reserve1).Float64()
				return f / 1e18
			}(),
		)),
	)
}
