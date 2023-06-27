package models

import "time"

type Battle struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type BattleDetail struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	PokemonId int       `gorm:"not null" json:"pokemon_id"`
	Score     int       `gorm:"not null" json:"score"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	BattleID  int
	Battle    *Battle
}
