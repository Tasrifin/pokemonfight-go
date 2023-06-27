package repositories

import "gorm.io/gorm"

type BattleRepo interface {
}

type battleRepo struct {
	db *gorm.DB
}

func NewBattleRepo(db *gorm.DB) BattleRepo {
	return &battleRepo{db}
}
