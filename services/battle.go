package services

import (
	"net/http"

	"github.com/Tasrifin/pokemonfight-go/constants"
	"github.com/Tasrifin/pokemonfight-go/helpers"
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
	for _, poke := range request.Pokemons {
		//CHECK POKEMON IS AVAILABLE
		url := constants.BASE_API_URL + "/" + constants.POKEMON_URI + "/" + poke.Name

		requestData, err := helpers.ReqHTTP(http.MethodGet, url)
		if err != nil {
			return &params.Response{
				Status: http.StatusInternalServerError,
				Payload: gin.H{
					"message": "error doing request API" + err.Error(),
				},
			}
		}

		if requestData.StatusCode != http.StatusOK {
			return &params.Response{
				Status: requestData.StatusCode,
				Payload: gin.H{
					"message": "Error on Pokemon : " + poke.Name,
				},
			}
		}

		defer requestData.Body.Close()

		//DOING BATTLE
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"data": "",
		},
	}
}
