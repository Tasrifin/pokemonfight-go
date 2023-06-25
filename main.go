package main

import (
	"github.com/Tasrifin/pokemonfight-go/config"
	"github.com/Tasrifin/pokemonfight-go/constants"
	"github.com/Tasrifin/pokemonfight-go/controllers"
	"github.com/Tasrifin/pokemonfight-go/repositories"
	"github.com/Tasrifin/pokemonfight-go/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	route := gin.Default()

	pokemonRepo := repositories.NewPokemonRepo(db)
	pokemonRepoService := services.NewPokemonService(pokemonRepo)
	pokemonController := controllers.NewPokemonController(pokemonRepoService)

	//ROUTE START
	route.GET("/pokemons", pokemonController.GetAllPokemons)

	route.Run(constants.APP_PORT)
}
