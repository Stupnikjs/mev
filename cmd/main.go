package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Stupnikjs/mev/internal/strategy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lmittmann/w3"
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
	ws_dRPC  = "wss://lb.drpc.live/ethereum/AhuxMhCqfkI8pF_0y4Fpi89GWcIMFIwR8ZsatuZZzRRv"
	pk       = os.Getenv("PRIVATE_KEY")
)

func main() {
	pool := strategy.V2Pool{
		Address: common.HexToAddress("0x0d4a11d5EEaac28EC3F61d100daF4d40471f1852"),
		Active:  true, // tu peux le confirmer dynamiquement comme dans l'exemple précédent
		Tokens: []string{
			"0xdAC17F958D2ee523a2206206994597C13D831ec7", // USDC
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH
		},
	}

	client, err := ethclient.Dial(rpc_pub)
	if err != nil {
		log.Fatal(err)
	}
	b, err := pool.IsActiveUniswapV2Pair(w3.NewClient(client.Client()))
	if err != nil {
		log.Fatal(err)
	}
	if b {
		fmt.Println("Active pool")
	}

}
