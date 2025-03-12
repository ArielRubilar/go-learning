package main

import (
	"fmt"

	"github.com/ArielRubilar/go-learning/application"
	"github.com/ArielRubilar/go-learning/infractucture/api"
	"github.com/ArielRubilar/go-learning/infractucture/store"
)

const (
	appName string = "Poke API"
)

func main() {
	fmt.Printf("Starting the application: %s\n", appName)

	repository := api.NewPokemonApi()

	store := store.NewCvsStore()

	service := application.NewPokemonService(repository, store)

	pokemon, err := service.GetByName("pikachu")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = service.SavePokemon(pokemon)

}
