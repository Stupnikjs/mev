package client

import (
	"crypto/ecdsa"

	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
)

type Clients struct {
	RPC   *w3.Client // connexion HTTP standard (simulation, eth_call)
	Relay *w3.Client // connexion relay Flashbots (envoi bundles)
}

func New(rpcURL, relayURL string, authKey *ecdsa.PrivateKey) (*Clients, error) {
	rpc := w3.MustDial(rpcURL)
	// eth_call → simuler des transactions
	// eth_getBlockByNumber → récupérer le bloc courant
	// eth_subscribe → écouter le mempool
	relay := flashbots.MustDial(relayURL, authKey)
	// AuthTransport injecte X-Flashbots-Signature automatiquement
	return &Clients{RPC: rpc, Relay: relay}, nil
}
