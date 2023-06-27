package models

type GetPokemonByAPI struct {
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Position int    `json:"position,omitempty"`
	Id       int    `json:"id"`
}

type GetTotalScores struct {
	PokemonId  int    `json:"pokemon_id"`
	Name       string `json:"name,omitempty"`
	TotalScore int    `json:"total_score"`
}
