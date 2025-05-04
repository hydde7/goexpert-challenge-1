package cep

import (
	"net/http"

	"github.com/hydde7/goexpert-challenge-1/internal/handler"
	"github.com/hydde7/goexpert-challenge-1/internal/services"
	"github.com/hydde7/goexpert-challenge-1/internal/utils"
)

type ControllerGetCepTemperature struct {
	handler.TransactionControllerImpl
}

func (c *ControllerGetCepTemperature) Execute(payload interface{}) (result handler.ResponseController) {
	result = handler.NewJsonResponseController()
	cep := c.GetParam("cep")

	if !utils.ValidateCEP(cep) {
		result.SetResult(http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	cepPayload, _ := services.ViaCepRequest(cep)
	if cepPayload.Cep == "" {
		result.SetResult(http.StatusNotFound, "zipcode not found")
		return
	}

	weatherPayload, err := services.FreeWeatherRequest(cepPayload.Localidade)
	if err != nil {
		result.SetResult(http.StatusInternalServerError, "internal server error")
		return
	}

	response := getCepTemperatureResponse{
		TemperatureCelsius:    weatherPayload.Current.TempC,
		TemperatureFahrenheit: weatherPayload.Current.TempF,
		TemperatureKelvin:     weatherPayload.Current.TempC + 273,
	}

	result.SetResult(http.StatusOK, response)
	return
}

type getCepTemperatureResponse struct {
	TemperatureCelsius    float64 `json:"temp_C"`
	TemperatureFahrenheit float64 `json:"temp_F"`
	TemperatureKelvin     float64 `json:"temp_K"`
}

func NewControllerGetCepTemperature() handler.Controller {
	return &ControllerGetCepTemperature{}
}
