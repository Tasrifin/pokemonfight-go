package controllers

import (
	"net/http"
	"strconv"

	"github.com/Tasrifin/pokemonfight-go/params"
	"github.com/Tasrifin/pokemonfight-go/services"
	"github.com/gin-gonic/gin"
)

type BattleController struct {
	battleService services.BattleService
}

func NewBattleController(service *services.BattleService) *BattleController {
	return &BattleController{
		battleService: *service,
	}
}

func (p *BattleController) CreateAutoBattle(c *gin.Context) {
	var req params.CreateAutoBattle

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := p.battleService.CreateAutoBattle(req)

	c.JSON(result.Status, result.Payload)
}

func (p *BattleController) GetTotalScores(c *gin.Context) {
	result := p.battleService.GetTotalScores()

	c.JSON(result.Status, result.Payload)
}

func (p *BattleController) BattleEliminatePokemon(c *gin.Context) {
	var req params.BattleEliminatePokemon

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	battleId, _ := strconv.Atoi(c.Param("battleID"))
	req.BattleID = battleId

	result := p.battleService.BattleEliminatePokemon(req)

	c.JSON(result.Status, result.Payload)
}
