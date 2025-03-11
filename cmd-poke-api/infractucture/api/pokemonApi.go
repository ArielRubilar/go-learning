package api

import (
	"encoding/json"
	"io"
	"net/http"

	model "github.com/ArielRubilar/go-learning/domain/models"
	repository "github.com/ArielRubilar/go-learning/domain/repositories"
)

var baseUri = "https://pokeapi.co/api/v2/pokemon"

type ApiPokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type ApiPokemon struct {
	Name   string           `json:"name"`
	Height int              `json:"height"`
	Weight int              `json:"weight"`
	Types  []ApiPokemonType `json:"types"`
}

type PokemonApi struct {
	uri string
}

func NewPokemonApi() repository.PokemonRepository {
	return &PokemonApi{
		uri: baseUri,
	}
}

func (api *PokemonApi) get(endpoint string) (data []byte, err error) {

	r, err := http.Get(api.uri + endpoint)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (api *PokemonApi) GetByName(name string) (p *model.Pokemon, err error) {
	endpoint := "/" + name

	data, err := api.get(endpoint)

	if err != nil {
		return nil, err
	}

	apiPokemon := ApiPokemon{}

	err = json.Unmarshal(data, &apiPokemon)

	if err != nil {

		return nil, err
	}

	var types []string

	for _, t := range apiPokemon.Types {
		types = append(types, t.Type.Name)
	}

	p = model.NewPokemon(apiPokemon.Name, apiPokemon.Height, apiPokemon.Weight, types)

	return
}
