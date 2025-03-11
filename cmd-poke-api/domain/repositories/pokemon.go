package repository

import (
	model "github.com/ArielRubilar/go-learning/domain/models"
)

type PokemonRepository interface {
	GetByName(name string) (*model.Pokemon, error)
}
