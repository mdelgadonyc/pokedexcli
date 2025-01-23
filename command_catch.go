package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemonResp.BaseXP)
	threshold := 40

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if res > threshold {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonResp.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
	return nil
}
