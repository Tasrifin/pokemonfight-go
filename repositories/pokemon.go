package repositories

import "gorm.io/gorm"

type PokemonRepo interface {
}

type pokemonRepo struct {
	db *gorm.DB
}

func NewPokemonRepo(db *gorm.DB) PokemonRepo {
	return &pokemonRepo{db}
}
