package config

import "os"

// Dans config.go

func Load() {
	_ = os.Getenv("BOT_PRIV_KEY")  // wallet avec des ETH → signe les txs
	_ = os.Getenv("AUTH_PRIV_KEY") // wallet vide → signe les headers HTTP
	/*

	   | Clé | Rôle | Besoin de fonds ? |
	   |-----|------|-------------------|
	   | `BotPrivKey` | Signe les transactions Ethereum dans le bundle | ✅ Oui (gas) |
	   | `AuthPrivKey` | Signe les headers HTTP pour s'authentifier au relay | ❌ Non |

	   C'est une séparation de sécurité : si ton `AuthPrivKey` leak, personne ne peut voler tes fonds.

	   ---
	*/
}
