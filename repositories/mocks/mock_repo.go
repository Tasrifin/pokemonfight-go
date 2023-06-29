package mocks

import (
	"time"

	"github.com/Tasrifin/pokemonfight-go/models"
	"github.com/stretchr/testify/mock"
)

type BattleRepository interface {
	CreateBattle(params models.Battle) (models.Battle, error)
	CreateBattleDetail(param []models.BattleDetail) (data []models.BattleDetail, err error)
	GetTotalScores() ([]models.GetTotalScores, error)
	GetBattleDetailByBattleID(battleId int) ([]models.BattleDetail, error)
	GetBattleDetailByIDAndPokemonID(battleId, pokemonId int) (models.BattleDetail, error)
	UpdateBattleDetailPokemon(data []models.BattleDetail) error
	GetAllBattleData(start, end time.Time) ([]models.Battle, error)
}

type MockBattleRepository struct {
	mock.Mock
}

func (_m *MockBattleRepository) CreateBattle(param models.Battle) (models.Battle, error) {
	args := _m.Called(param)
	return args.Get(0).(models.Battle), nil
}

func (_m *MockBattleRepository) CreateBattleDetail(param []models.BattleDetail) (data []models.BattleDetail, err error) {
	args := _m.Called(param)
	return args.Get(0).([]models.BattleDetail), nil
}

func (_m *MockBattleRepository) GetTotalScores() ([]models.GetTotalScores, error) {
	args := _m.Called()
	return args.Get(0).([]models.GetTotalScores), nil
}

func (_m *MockBattleRepository) GetBattleDetailByBattleID(battleId int) ([]models.BattleDetail, error) {
	args := _m.Called(battleId)
	return args.Get(0).([]models.BattleDetail), nil
}

func (_m *MockBattleRepository) GetBattleDetailByIDAndPokemonID(battleId, pokemonId int) (models.BattleDetail, error) {
	args := _m.Called(battleId, pokemonId)
	r0, _ := args[0].(models.BattleDetail)
	r1, _ := args[1].(error)
	return r0, r1
}

func (_m *MockBattleRepository) UpdateBattleDetailPokemon(data []models.BattleDetail) error {
	args := _m.Called(data)
	r0, _ := args[0].(error)
	return r0
}

func (_m *MockBattleRepository) GetAllBattleData(start, end time.Time) ([]models.Battle, error) {
	args := _m.Called(start, end)
	return args.Get(0).([]models.Battle), nil
}
