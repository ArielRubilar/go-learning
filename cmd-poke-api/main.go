package main

import (
	"fmt"

	"github.com/ArielRubilar/go-learning/application"
	"github.com/ArielRubilar/go-learning/infractucture/api"
)

const (
	appName string = "Poke API"
)

func main() {
	fmt.Printf("Starting the application: %s\n", appName)

	repository := api.NewPokemonApi()

	service := application.NewPokemonService(repository)

	pokemon, err := service.GetPokemon("pikachu")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("Pokemon: %s\n", pokemon.String())

}
