package strategy

import (
	"github.com/lmittmann/w3"
)

var UniswapV2SwapSelectors = map[string]string{
	"0x38ed1739": "swapExactTokensForTokens",
	"0x8803dbee": "swapTokensForExactTokens",
	"0x7ff36ab5": "swapExactETHForTokens",
	"0xfb3bdb41": "swapETHForExactTokens",
	"0x791ac947": "swapExactTokensForETH",
	"0x18cbafe5": "swapTokensForExactETH",
	"0x5ae401dc": "multicall", // souvent utilisé par des routers/aggregators
}

var UniswapV2Funcs = map[string]*w3.Func{
	"0x38ed1739": w3.MustNewFunc("swapExactTokensForTokens(uint256,uint256,address[],address,uint256)", ""),
	"0x8803dbee": w3.MustNewFunc("swapTokensForExactTokens(uint256,uint256,address[],address,uint256)", ""),
	"0x7ff36ab5": w3.MustNewFunc("swapExactETHForTokens(uint256,address[],address,uint256)", ""),
	"0xfb3bdb41": w3.MustNewFunc("swapETHForExactTokens(uint256,address[],address,uint256)", ""),
	"0x791ac947": w3.MustNewFunc("swapExactTokensForETH(uint256,uint256,address[],address,uint256)", ""),
	"0x18cbafe5": w3.MustNewFunc("swapTokensForExactETH(uint256,uint256,address[],address,uint256)", ""),
	// ajoute les autres si besoin (addLiquidity, removeLiquidity...)
}
