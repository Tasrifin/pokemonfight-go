package services

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/Tasrifin/pokemonfight-go/constants"
	"github.com/Tasrifin/pokemonfight-go/models"
	"github.com/Tasrifin/pokemonfight-go/params"
	"github.com/Tasrifin/pokemonfight-go/repositories/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTotalScores(t *testing.T) {
	repoMock := new(mocks.MockBattleRepository)

	data := []models.GetTotalScores{
		{
			PokemonId:  1,
			TotalScore: 5,
		},
		{
			PokemonId:  4,
			TotalScore: 2,
		},
		{
			PokemonId:  5,
			TotalScore: 1,
		},
		{
			PokemonId:  3,
			TotalScore: 4,
		},
		{
			PokemonId:  2,
			TotalScore: 3,
		},
	}

	listTests := []struct {
		name      string
		mock      func()
		want      *params.Response
		isSuccess bool
	}{
		{
			name: "Get Total Score",
			mock: func() {
				repoMock.On("GetTotalScores", mock.Anything).Return(data, nil).Once()
			},
			want: &params.Response{
				Status: http.StatusOK,
				Payload: gin.H{
					"data": data,
				},
			},
			isSuccess: true,
		},
		{
			name: "Get Total Score Fail",
			mock: func() {
				repoMock.On("GetTotalScores", mock.Anything).Return([]models.GetTotalScores{}, errors.New("db error")).Once()
			},
			want: &params.Response{
				Status: http.StatusInternalServerError,
				Payload: gin.H{
					"error": "db error",
				},
			},
			isSuccess: false,
		},
	}

	for _, test := range listTests {
		test.mock()
		battleService := BattleService{battleRepo: repoMock}
		res := battleService.GetTotalScores()

		assert.Equal(t, test.want, res)
	}
}

func TestGelAllBattleData(t *testing.T) {
	repoMock := new(mocks.MockBattleRepository)

	battleDetail := []models.BattleDetail{
		{
			ID:        1,
			PokemonId: 1,
			Score:     3,
		},
		{
			ID:        2,
			PokemonId: 4,
			Score:     5,
		},
		{
			ID:        3,
			PokemonId: 5,
			Score:     2,
		},
		{
			ID:        4,
			PokemonId: 2,
			Score:     1,
		},
		{
			ID:        5,
			PokemonId: 3,
			Score:     4,
		},
	}
	data := []models.Battle{
		{
			ID:            1,
			Name:          "Battle 1",
			BattleDetails: battleDetail,
		},
	}
	startString := "2023-06-28 00:00:00"
	endString := "2023-06-30 00:00:00"
	startDate, _ := time.Parse(constants.DATETIME_LAYOUT, startString)
	endDate, _ := time.Parse(constants.DATETIME_LAYOUT, endString)

	listTests := []struct {
		name      string
		mock      func()
		request   params.GetBattleData
		want      *params.Response
		isSuccess bool
	}{
		{
			name: "Get All Battle",
			request: params.GetBattleData{
				StartDate: startString,
				EndDate:   endString,
			},
			mock: func() {
				repoMock.On("GetBattleDetailByBattleID", 1).Return(battleDetail, nil).Once()
				repoMock.On("GetAllBattleData", startDate, endDate).Return(data, nil).Once()
			},
			want: &params.Response{
				Status: http.StatusOK,
				Payload: gin.H{
					"data": data,
				},
			},
			isSuccess: true,
		},
		{
			name: "Get All Battle Fail",
			request: params.GetBattleData{
				StartDate: endString,
				EndDate:   startString,
			},
			mock: func() {
				repoMock.On("GetBattleDetailByBattleID", 1).Return([]models.BattleDetail{}, errors.New("error")).Once()
				repoMock.On("GetAllBattleData", startDate, endDate).Return([]models.Battle{}, errors.New("error")).Once()
			},
			want: &params.Response{
				Status: http.StatusBadRequest,
				Payload: gin.H{
					"error": "end date must be greater than start date",
				},
			},
			isSuccess: false,
		},
	}

	for _, test := range listTests {
		test.mock()
		battleService := BattleService{battleRepo: repoMock}
		res := battleService.GetAllBattleData(test.request)

		assert.Equal(t, test.want, res)
	}
}

func TestCreateBattle(t *testing.T) {
	repoMock := new(mocks.MockBattleRepository)

	data := models.Battle{
		Name: "Battle 1",
	}

	battleDetail := []models.BattleDetail{
		{
			PokemonId: 1,
			Score:     2,
		},
		{
			PokemonId: 2,
			Score:     5,
		},
		{
			PokemonId: 3,
			Score:     4,
		},
		{
			PokemonId: 4,
			Score:     3,
		},
		{
			PokemonId: 5,
			Score:     1,
		},
	}

	listTests := []struct {
		name      string
		mock      func()
		request   params.CreateAutoBattle
		want      *params.Response
		isSuccess bool
	}{
		{
			name: "Create Battle",
			request: params.CreateAutoBattle{
				BattleName: "Battle 1",
				Pokemons:   []int{1, 2, 3, 4, 5},
			},
			mock: func() {
				repoMock.On("CreateBattle", data).Return(data, nil).Once()
				repoMock.On("CreateBattleDetail", battleDetail).Return(battleDetail, nil).Once()
			},
			want: &params.Response{
				Status: http.StatusOK,
				Payload: gin.H{
					"message": "success",
				},
			},
			isSuccess: true,
		},
		{
			name: "Create Battle Fail",
			request: params.CreateAutoBattle{
				BattleName: "Battle 2",
				Pokemons:   []int{1, 2, 3, 4},
			},
			mock: func() {
				repoMock.On("CreateBattle", data).Return(models.Battle{}, errors.New("error")).Once()
				repoMock.On("CreateBattleDetail", battleDetail).Return([]models.BattleDetail{}, errors.New("error")).Once()
			},
			want: &params.Response{
				Status: http.StatusBadRequest,
				Payload: gin.H{
					"error": "total Pokemons must be 5",
				},
			},
			isSuccess: false,
		},
	}

	for _, test := range listTests {
		test.mock()
		battleService := BattleService{battleRepo: repoMock}
		res := battleService.CreateAutoBattle(test.request)

		assert.Equal(t, res, test.want)
	}
}

func TestEliminatePokemon(t *testing.T) {
	repoMock := new(mocks.MockBattleRepository)

	battleDetailSingle := models.BattleDetail{ID: 3,
		PokemonId: 3,
		Score:     4,
		BattleID:  1,
	}
	battleDetail := []models.BattleDetail{
		{
			ID:        1,
			PokemonId: 1,
			Score:     2,
			BattleID:  1,
		},
		{
			ID:        2,
			PokemonId: 2,
			Score:     5,
			BattleID:  1,
		},
		{
			ID:        3,
			PokemonId: 3,
			Score:     4,
			BattleID:  1,
		},
		{
			ID:        4,
			PokemonId: 4,
			Score:     3,
			BattleID:  1,
		},
		{
			ID:        5,
			PokemonId: 5,
			Score:     1,
			BattleID:  1,
		},
	}
	updateData := []models.BattleDetail{
		{
			ID:        1,
			PokemonId: 1,
			Score:     3,
			BattleID:  1,
		},
		{
			ID:        4,
			PokemonId: 4,
			Score:     4,
			BattleID:  1,
		},
		{
			ID:        5,
			PokemonId: 5,
			Score:     2,
			BattleID:  1,
		},
		{
			ID:        3,
			PokemonId: 3,
			Score:     0,
			BattleID:  1,
		},
	}

	listTests := []struct {
		name      string
		mock      func()
		request   params.BattleEliminatePokemon
		want      *params.Response
		isSuccess bool
	}{
		{
			name: "Eliminate Pokemon",
			request: params.BattleEliminatePokemon{
				BattleID:  1,
				PokemonID: 3,
			},
			mock: func() {
				repoMock.On("GetBattleDetailByIDAndPokemonID", 1, 3).Return(battleDetailSingle, nil).Once()
				repoMock.On("GetBattleDetailByBattleID", 1).Return(battleDetail, nil).Once()
				repoMock.On("UpdateBattleDetailPokemon", updateData).Return(nil).Once()
			},
			want: &params.Response{
				Status: http.StatusOK,
				Payload: gin.H{
					"message": "success",
				},
			},
			isSuccess: true,
		},
		{
			name: "Eliminate Pokemon Fail",
			request: params.BattleEliminatePokemon{
				BattleID:  0,
				PokemonID: 0,
			},
			mock: func() {
				repoMock.On("GetBattleDetailByIDAndPokemonID", 0, 0).Return(models.BattleDetail{}, nil).Once()
				repoMock.On("GetBattleDetailByBattleID", 1).Return([]models.BattleDetail{}, errors.New("error")).Once()
				repoMock.On("UpdateBattleDetailPokemon", updateData).Return(errors.New("error")).Once()
			},
			want: &params.Response{Status: http.StatusNotFound,
				Payload: gin.H{
					"error": "detail not found",
				}},
			isSuccess: false,
		},
	}

	for _, test := range listTests {
		test.mock()
		battleService := BattleService{battleRepo: repoMock}
		res := battleService.BattleEliminatePokemon(test.request)

		assert.Equal(t, res, test.want)
	}
}
