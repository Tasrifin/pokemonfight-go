package repositories

import (
	"github.com/Tasrifin/pokemonfight-go/models"
	"gorm.io/gorm"
)

type BattleRepo interface {
	CreateBattle(data *models.Battle) (*models.Battle, error)
	CreateBattleDetail(data *models.BattleDetail) (*models.BattleDetail, error)
	GetTotalScores() *[]models.GetTotalScores
}

type battleRepo struct {
	db *gorm.DB
}

func NewBattleRepo(db *gorm.DB) BattleRepo {
	return &battleRepo{db}
}

func (b *battleRepo) CreateBattle(data *models.Battle) (*models.Battle, error) {
	return data, b.db.Create(&data).Error
}

func (b *battleRepo) CreateBattleDetail(data *models.BattleDetail) (*models.BattleDetail, error) {
	return data, b.db.Create(&data).Error
}

func (b *battleRepo) GetTotalScores() *[]models.GetTotalScores {
	var totalScores []models.GetTotalScores
	_ = b.db.Table("battle_details").Select("pokemon_id, sum(score) as total_score").Group("pokemon_id").Order("total_score desc").Scan(&totalScores)
	return &totalScores
}
