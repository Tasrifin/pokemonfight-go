package params

type CreateAutoPokemon struct {
	ID int `json:"id" binding:"required"`
}

type CreateManualPokemon struct {
	ID       int `json:"id" binding:"required"`
	Position int `json:"position" binding:"required"`
}
