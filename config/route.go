package config

import (
	"os"

	"github.com/Tasrifin/pokemonfight-go/controllers"
	"github.com/Tasrifin/pokemonfight-go/repositories"
	"github.com/Tasrifin/pokemonfight-go/services"
	"github.com/gin-gonic/gin"
)

func InitRoute() {
	db := ConnectDB()
	route := gin.Default()

	//POKE REPO
	pokemonRepo := repositories.NewPokemonRepo(db)
	pokemonService := services.NewPokemonService(pokemonRepo)
	pokemonController := controllers.NewPokemonController(pokemonService)

	//BATTLE REPO
	battleRepo := repositories.NewBattleRepo(db)
	battleService := services.NewBattleService(battleRepo)
	battleController := controllers.NewBattleController(battleService)

	//ROUTE START
	route.GET("/pokemons", pokemonController.GetAllPokemons)
	route.GET("/pokemon/scores", battleController.GetTotalScores)
	route.GET("/battles", battleController.GetAllBattleData)
	route.POST("/battle/auto", battleController.CreateAutoBattle)
	route.PUT("/battle/:battleID/eliminate-pokemon", battleController.BattleEliminatePokemon)

	route.Run(os.Getenv("APP_PORT"))
}
