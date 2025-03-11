package application

import (
	model "github.com/ArielRubilar/go-learning/domain/models"
	repository "github.com/ArielRubilar/go-learning/domain/repositories"
)

type PokemonService struct {
	repository repository.PokemonRepository
}

func NewPokemonService(r repository.PokemonRepository) *PokemonService {
	return &PokemonService{
		repository: r,
	}
}

func (s *PokemonService) GetPokemon(name string) (*model.Pokemon, error) {

	pokemon, err := s.repository.GetByName(name)

	if err != nil {
		return nil, err
	}

	return &model.Pokemon{
		Name:   pokemon.Name,
		Height: pokemon.Height,
		Weight: pokemon.Weight,
		Types:  pokemon.Types,
	}, nil
}
