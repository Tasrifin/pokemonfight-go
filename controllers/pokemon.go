package controllers

import (
	"github.com/Tasrifin/pokemonfight-go/services"
	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	pokemonService services.PokemonService
}

func NewPokemonController(service *services.PokemonService) *PokemonController {
	return &PokemonController{
		pokemonService: *service,
	}
}

func (p *PokemonController) GetAllPokemons(c *gin.Context) {
	result := p.pokemonService.GetAllPokemons()

	c.JSON(result.Status, result.Payload)
}
