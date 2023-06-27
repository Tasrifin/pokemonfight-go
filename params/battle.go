package params

type CreateAutoBattle struct {
	BattleName string `json:"battle_name" binding:"required"`
	Pokemons   []int  `json:"pokemons" binding:"required"`
}

type CreateManualBattle struct {
	BattleName string `json:"battle_name" binding:"required"`
	Pokemons   []int  `json:"pokemons" binding:"required"`
}

type BattleEliminatePokemon struct {
	BattleID  int `json:"battle_id"`
	PokemonID int `json:"pokemon_id" binding:"required"`
}
