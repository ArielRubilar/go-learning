package application

import (
	"fmt"

	model "github.com/ArielRubilar/go-learning/domain/models"
	repository "github.com/ArielRubilar/go-learning/domain/repositories"
)

type PokemonService struct {
	repository repository.PokemonRepository
	store      repository.Store
}

func NewPokemonService(r repository.PokemonRepository, s repository.Store) *PokemonService {
	return &PokemonService{
		repository: r,
		store:      s,
	}
}

func (s *PokemonService) GetByName(name string) (*model.Pokemon, error) {

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

func (s *PokemonService) SavePokemon(pokemon *model.Pokemon) error {
	data := make(map[string]string)

	data["Name"] = pokemon.Name
	data["Height"] = fmt.Sprintf("%d", pokemon.Height)
	data["Weight"] = fmt.Sprintf("%d", pokemon.Weight)
	data["Types"] = pokemon.Types.String()

	err := s.store.Save("data/pokemon-"+pokemon.Name, []map[string]string{data})

	if err != nil {
		return err
	}

	return nil
}
