package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tasrifin/pokemonfight-go/constants"
	"github.com/Tasrifin/pokemonfight-go/helpers"
	"github.com/Tasrifin/pokemonfight-go/models"
	"github.com/Tasrifin/pokemonfight-go/params"
	"github.com/Tasrifin/pokemonfight-go/repositories"
	"github.com/gin-gonic/gin"
)

type PokemonService struct {
	pokemonRepo repositories.PokemonRepo
}

func NewPokemonService(pokeRepo repositories.PokemonRepo) *PokemonService {
	return &PokemonService{
		pokemonRepo: pokeRepo,
	}
}

func (p *PokemonService) GetAllPokemons() *params.Response {
	url := constants.BASE_API_URL + "/" + constants.POKEMON_URI + "?limit=" + fmt.Sprint(constants.DEFAULT_LIMIT)

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
		}
	}

	defer requestData.Body.Close()

	var resultData models.GetPokemonByAPI

	err = json.NewDecoder(requestData.Body).Decode(&resultData)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	for i, data := range resultData.Results {
		id, err := helpers.GetIDByUrl(data.Url)
		if err != nil {
			return &params.Response{
				Status: http.StatusInternalServerError,
				Payload: gin.H{
					"error": err.Error(),
				},
			}
		}

		resultData.Results[i].Id = id
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"data": resultData.Results,
		},
	}
}
