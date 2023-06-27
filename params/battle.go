package params

type CreateAutoBattle struct {
	BattleName string `json:"battle_name" binding:"required"`
	Pokemons   []int  `json:"pokemons" binding:"required"`
}

type CreateManualBattle struct {
	BattleName string `json:"battle_name" binding:"required"`
	Pokemons   []int  `json:"pokemons" binding:"required"`
}
