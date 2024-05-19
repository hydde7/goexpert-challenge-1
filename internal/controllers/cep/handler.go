package cep

import (
	"github.com/gin-gonic/gin"
	"github.com/hydde7/goexpert-challenge-1/internal/handler"
)

// GET Endpoints

// [GET] /cep/{cep}
// @Summary Get temperature by zipcode
// @Tags CEP
// @Description Get temperature by zipcode
// @Accept json
// @Produce json
// @Param cep path string true "CEP"
// @Success 200 {object} getCepTemperatureResponse "Temperature"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /cep/{cep} [get]
func HandleGetCepTemperature(c *gin.Context) {
	handler.RequestWithController(c, nil, NewControllerGetCepTemperature())
}
