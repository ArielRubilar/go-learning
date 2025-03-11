package model

import "strconv"

type PokemonType string

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Types  []PokemonType
}

func (p *Pokemon) String() string {
	// json string
	types := ""
	for i, t := range p.Types {
		if i == 0 {
			types += `"` + t.String() + `"`
		} else {
			types += `,"` + t.String() + `"`
		}
	}
	return `{"name":"` + p.Name + `","height":` + strconv.Itoa(p.Height) + `,"weight":` + strconv.Itoa(p.Weight) + `,"types":[` + types + `]}`
}

func (t *PokemonType) String() string {
	return string(*t)
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
