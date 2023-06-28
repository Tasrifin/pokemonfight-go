package repositories

import (
	"time"

	"github.com/Tasrifin/pokemonfight-go/models"
	"gorm.io/gorm"
)

type BattleRepo interface {
	CreateBattle(data *models.Battle) (*models.Battle, error)
	CreateBattleDetail(data *models.BattleDetail) (*models.BattleDetail, error)
	GetTotalScores() ([]models.GetTotalScores, error)
	GetBattleDetailByBattleID(battleId int) ([]models.BattleDetail, error)
	GetBattleDetailByIDAndPokemonID(battleId, pokemonId int) (models.BattleDetail, error)
	UpdateBattleDetailPokemon(detailId int, updateData models.BattleDetail) error
	GetAllBattleData(start, end time.Time) ([]models.Battle, error)
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

func (b *battleRepo) GetTotalScores() ([]models.GetTotalScores, error) {
	var totalScores []models.GetTotalScores
	err := b.db.Table("battle_details").Select("pokemon_id, sum(score) as total_score").Group("pokemon_id").Order("total_score desc").Scan(&totalScores).Error
	return totalScores, err
}

func (b *battleRepo) GetBattleDetailByBattleID(battleId int) ([]models.BattleDetail, error) {
	var detail []models.BattleDetail
	err := b.db.Table("battle_details").Where("battle_id=?", battleId).Order("score desc").Scan(&detail).Error
	return detail, err
}

func (b *battleRepo) GetBattleDetailByIDAndPokemonID(battleId, pokemonId int) (models.BattleDetail, error) {
	var detail models.BattleDetail
	err := b.db.Table("battle_details").Where("pokemon_id=?", pokemonId).Where("battle_id=?", battleId).Scan(&detail).Error
	return detail, err
}

func (b *battleRepo) UpdateBattleDetailPokemon(detailId int, updateData models.BattleDetail) error {
	err := b.db.Table("battle_details").Select("score", "updated_at").Where("id=?", detailId).Updates(updateData).Error
	return err
}

func (b *battleRepo) GetAllBattleData(start, end time.Time) ([]models.Battle, error) {
	var battles []models.Battle

	if start.IsZero() && end.IsZero() {
		err := b.db.Table("battles").Order("id").Scan(&battles).Error
		return battles, err
	}

	err := b.db.Table("battles").Where("created_at between ? and ?", start, end).Order("id").Scan(&battles).Error
	return battles, err
}
