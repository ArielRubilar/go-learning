package api

import (
	model "github.com/ArielRubilar/go-learning/domain/models"
	repository "github.com/ArielRubilar/go-learning/domain/repositories"
)

var baseUri = "https://pokeapi.co/api/v2/pokemon"

type PokemonApi struct {
	uri string
}

func NewPokemonApi() repository.PokemonRepository {
	return &PokemonApi{
		uri: baseUri,
	}
}

func (p *PokemonApi) GetByName(name string) (*model.Pokemon, error) {
	return &model.Pokemon{
		Name:   "pikachu",
		Height: 4,
		Weight: 60,
		Types:  []model.Type{"electric"},
	}, nil
}
