package params

type CreateAutoPokemon struct {
	Name string `json:"name" binding:"required"`
}

type CreateManualPokemon struct {
	Name     string `json:"name" binding:"required"`
	Position int    `json:"position" binding:"required"`
}
