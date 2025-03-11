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

	service := application.NewPokemonService(repository)

	pokemon, err := service.GetPokemon("bulbasaur")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	store := store.NewCvsStore()

	data := make(map[string]string)

	data["Name"] = pokemon.Name
	data["Height"] = fmt.Sprintf("%d", pokemon.Height)
	data["Weight"] = fmt.Sprintf("%d", pokemon.Weight)
	data["Types"] = pokemon.Types.String()

	_ = store.Save("pokemon", []map[string]string{data})

	fmt.Printf("Pokemon: %s\n", pokemon.String())

}
