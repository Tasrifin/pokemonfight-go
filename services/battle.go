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
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": "Pokemon is duplicated, please re-check",
			},
		}
	}

	if len(request.Pokemons) != 5 {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": "Total Pokemons must be 5",
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
		Name:      request.BattleName,
		CreatedAt: time.Now(),
	}

	savedBattle, err := b.battleRepo.CreateBattle(&battleMaster)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": "While create battle - " + err.Error(),
			},
		}
	}

	for _, detail := range battleDetails {
		detailData := models.BattleDetail{
			PokemonId: detail.PokemonId,
			Score:     detail.Score,
			BattleID:  savedBattle.ID,
			CreatedAt: time.Now(),
		}

		_, err := b.battleRepo.CreateBattleDetail(&detailData)
		if err != nil {
			return &params.Response{
				Status: http.StatusInternalServerError,
				Payload: gin.H{
					"error": "While create detail battle - " + err.Error(),
				},
			}
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "Success",
		},
	}
}

func (b *BattleService) GetTotalScores() *params.Response {
	data := b.battleRepo.GetTotalScores()

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"data": data,
		},
	}
}
