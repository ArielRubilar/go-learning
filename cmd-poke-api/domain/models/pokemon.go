package model

import "strconv"

type PokemonType string

type PokemonTypes []PokemonType

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Types  PokemonTypes
}

func (p *Pokemon) String() string {
	return `{"name":"` + p.Name + `","height":` + strconv.Itoa(p.Height) + `,"weight":` + strconv.Itoa(p.Weight) + `,"types":` + p.Types.String() + `}`
}

func (t *PokemonType) String() string {
	return string(*t)
}

func (ts *PokemonTypes) String() string {
	types := ""
	for i, t := range *ts {
		if i == 0 {
			types += `"` + t.String() + `"`
		} else {
			types += `,"` + t.String() + `"`
		}
	}
	return `[` + types + `]`
}

func toPokemonTypes(types []string) (pokemonTypes []PokemonType) {
	for _, t := range types {
		pokemonTypes = append(pokemonTypes, PokemonType(t))
	}
	return pokemonTypes
}

func NewPokemon(name string, height int, weight int, types []string) *Pokemon {
	return &Pokemon{
		Name:   name,
		Height: height,
		Weight: weight,
		Types:  toPokemonTypes(types),
	}
}
