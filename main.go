package main

import (
	"time"

	"github.com/mdelgadonyc/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: make(map[string]pokeapi.RespPokemon),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
