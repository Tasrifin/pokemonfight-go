package models

type GetPokemonByAPI struct {
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Position int    `json:"position,omitempty"`
}
