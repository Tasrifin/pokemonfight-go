package services

import (
	"errors"
	"log"
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
		name    string
		mock    func()
		want    *params.Response
		wantErr bool
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
			wantErr: false,
		},
		{
			name: "Get Total Score Fail",
			mock: func() {
				repoMock.On("GetTotalScores", mock.Anything).Return([]models.GetTotalScores{}, errors.New("error")).Once()
			},
			want:    &params.Response{},
			wantErr: true,
		},
	}

	for _, test := range listTests {
		test.mock()
		battleService := BattleService{battleRepo: repoMock}
		res := battleService.GetTotalScores()

		if !test.wantErr {
			log.Print(test.name)
			assert.Equal(t, test.want, res)
		}
		if test.wantErr {
			log.Print(test.name)
			assert.Error(t, errors.New("error"), res)
		}
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

	request := params.GetBattleData{
		StartDate: startString,
		EndDate:   endString,
	}

	listTests := []struct {
		name    string
		mock    func()
		want    *params.Response
		wantErr bool
	}{
		{
			name: "Get All Battle",
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
			wantErr: false,
		},
		{
			name: "Get All Battle Fail",
			mock: func() {
				repoMock.On("GetBattleDetailByBattleID", 1).Return([]models.BattleDetail{}, errors.New("error")).Once()
				repoMock.On("GetAllBattleData", startDate, endDate).Return([]models.Battle{}, errors.New("error")).Once()
			},
			want:    &params.Response{},
			wantErr: true,
		},
	}

	for _, test := range listTests {
		test.mock()
		battleService := BattleService{battleRepo: repoMock}
		res := battleService.GetAllBattleData(request)

		if !test.wantErr {
			log.Print(test.name)
			assert.Equal(t, test.want, res)
		}
		if test.wantErr {
			log.Print(test.name)
			assert.Error(t, errors.New("error"), res)
		}
	}
}

func TestCreateBattle(t *testing.T) {
	repoMock := new(mocks.MockBattleRepository)

	request := params.CreateAutoBattle{
		BattleName: "Battle 1",
		Pokemons:   []int{1, 2, 3, 4, 5},
	}

	data := models.Battle{
		Name: request.BattleName,
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
		name    string
		mock    func()
		args    params.CreateAutoBattle
		want    *params.Response
		wantErr bool
	}{
		{
			name: "Create Battle",
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
			wantErr: false,
		},
		// {
		// 	name: "Get Battle Fail",
		// 	mock: func() {
		// 		repoMock.On("CreateBattle", data).Return(models.Battle{}, errors.New("error")).Once()
		// 		repoMock.On("CreateBattleDetail", battleDetail).Return([]models.BattleDetail{}, errors.New("error")).Once()
		// 	},
		// 	want:    &params.Response{},
		// 	wantErr: true,
		// },
	}

	for _, test := range listTests {
		test.mock()
		battleService := BattleService{battleRepo: repoMock}
		res := battleService.CreateAutoBattle(request)

		if !test.wantErr {
			assert.Equal(t, res, test.want)
		}
		if test.wantErr {
			log.Print(test.name)
			assert.Error(t, errors.New("error"), res)
		}

	}
}

func TestEliminatePokemon(t *testing.T) {
	repoMock := new(mocks.MockBattleRepository)

	request := params.BattleEliminatePokemon{
		BattleID:  1,
		PokemonID: 3,
	}
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
		name    string
		mock    func()
		args    params.CreateAutoBattle
		want    *params.Response
		wantErr bool
	}{
		{
			name: "Eliminate Pokemon",
			mock: func() {
				repoMock.On("GetBattleDetailByIDAndPokemonID", request.BattleID, request.PokemonID).Return(battleDetailSingle, nil).Once()
				repoMock.On("GetBattleDetailByBattleID", 1).Return(battleDetail, nil).Once()
				repoMock.On("UpdateBattleDetailPokemon", updateData).Return(nil).Once()
			},
			want: &params.Response{
				Status: http.StatusOK,
				Payload: gin.H{
					"message": "success",
				},
			},
			wantErr: false,
		},
		{
			name: "Eliminate Pokemon Fail",
			mock: func() {
				repoMock.On("GetBattleDetailByIDAndPokemonID", request.BattleID, request.PokemonID).Return(models.BattleDetail{}, errors.New("error")).Once()
				repoMock.On("GetBattleDetailByBattleID", 1).Return([]models.BattleDetail{}, errors.New("error")).Once()
				repoMock.On("UpdateBattleDetailPokemon", updateData).Return(errors.New("error")).Once()
			},
			want:    &params.Response{},
			wantErr: true,
		},
	}

	for _, test := range listTests {
		test.mock()
		battleService := BattleService{battleRepo: repoMock}
		res := battleService.BattleEliminatePokemon(request)

		if !test.wantErr {
			assert.Equal(t, res, test.want)
		}
		if test.wantErr {
			log.Print(test.name)
			assert.Error(t, errors.New("error"), res)
		}

	}
}
