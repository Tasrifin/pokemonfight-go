package services

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Tasrifin/pokemonfight-go/constants"
	"github.com/Tasrifin/pokemonfight-go/helpers"
	"github.com/Tasrifin/pokemonfight-go/models"
	"github.com/Tasrifin/pokemonfight-go/params"
	"github.com/Tasrifin/pokemonfight-go/repositories"
	"github.com/gin-gonic/gin"
)

type BattleService struct {
	battleRepo repositories.BattleRepo
}

func NewBattleService(battleRepo repositories.BattleRepo) *BattleService {
	return &BattleService{
		battleRepo: battleRepo,
	}
}

func (b *BattleService) CreateAutoBattle(request params.CreateAutoBattle) *params.Response {
	battleDetails := []models.BattleDetail{}
	scores := []int{1, 2, 3, 4, 5}

	checkDuplicate := helpers.CheckDuplicateID(request.Pokemons)
	if checkDuplicate {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": "pokemon is duplicated, please re-check",
			},
		}
	}

	if len(request.Pokemons) != 5 {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": "total Pokemons must be 5",
			},
		}
	}

	for _, poke := range request.Pokemons {
		url := constants.BASE_API_URL + "/" + constants.POKEMON_URI + "/" + fmt.Sprint(poke)

		requestData, err := helpers.ReqHTTP(http.MethodGet, url)
		if err != nil {
			return &params.Response{
				Status: http.StatusInternalServerError,
				Payload: gin.H{
					"error": "error doing request API" + err.Error(),
				},
			}
		}

		if requestData.StatusCode != http.StatusOK {
			return &params.Response{
				Status: requestData.StatusCode,
				Payload: gin.H{
					"error": fmt.Sprint("Error on Pokemon : ", poke),
				},
			}
		}

		defer requestData.Body.Close()

		randomScore := rand.Intn(len(scores))
		score := scores[randomScore]
		scores = append(scores[:randomScore], scores[randomScore+1:]...)

		pokemonDetail := models.BattleDetail{
			PokemonId: poke,
			Score:     score,
		}

		battleDetails = append(battleDetails, pokemonDetail)
	}

	battleMaster := models.Battle{
		Name: request.BattleName,
	}

	createdBattle, err := b.battleRepo.CreateBattle(battleMaster)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": "While create battle - " + err.Error(),
			},
		}
	}

	for i := range battleDetails {
		battleDetails[i].BattleID = createdBattle.ID
	}
	_, err = b.battleRepo.CreateBattleDetail(battleDetails)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": "While create detail battle - " + err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "success",
		},
	}
}

func (b *BattleService) GetTotalScores() *params.Response {
	data, err := b.battleRepo.GetTotalScores()
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"data": data,
		},
	}
}

func (b *BattleService) BattleEliminatePokemon(request params.BattleEliminatePokemon) *params.Response {
	detailPokemonId, err := b.battleRepo.GetBattleDetailByIDAndPokemonID(request.BattleID, request.PokemonID)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}
	if detailPokemonId.ID == 0 {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "detail not found",
			},
		}
	}

	battleDetails, err := b.battleRepo.GetBattleDetailByBattleID(request.BattleID)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}
	if len(battleDetails) == 0 {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "detail not found",
			},
		}
	}

	updateDetails := []models.BattleDetail{}
	for _, detail := range battleDetails {
		if detail.Score < detailPokemonId.Score {
			updateData := models.BattleDetail{
				ID:        detail.ID,
				PokemonId: detail.PokemonId,
				Score:     detail.Score + 1,
				BattleID:  detail.BattleID,
			}
			updateDetails = append(updateDetails, updateData)
		}
	}

	updateDetails = append(updateDetails, models.BattleDetail{
		ID:        detailPokemonId.ID,
		PokemonId: detailPokemonId.PokemonId,
		Score:     0,
		BattleID:  detailPokemonId.BattleID,
	})

	err = b.battleRepo.UpdateBattleDetailPokemon(updateDetails)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "success",
		},
	}

}

func (b *BattleService) GetAllBattleData(request params.GetBattleData) *params.Response {
	startDate, _ := time.Parse(constants.DATETIME_LAYOUT, request.StartDate)
	endDate, _ := time.Parse(constants.DATETIME_LAYOUT, request.EndDate)

	if startDate.After(endDate) {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": "end date must be greater than start date",
			},
		}
	}

	battleData, err := b.battleRepo.GetAllBattleData(startDate, endDate)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	for i, battle := range battleData {
		battleDetails, err := b.battleRepo.GetBattleDetailByBattleID(battle.ID)
		if err != nil {
			return &params.Response{
				Status: http.StatusInternalServerError,
				Payload: gin.H{
					"error": err.Error(),
				},
			}
		}

		battleData[i].BattleDetails = battleDetails
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"data": battleData,
		},
	}
}
