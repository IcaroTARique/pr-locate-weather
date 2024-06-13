package main

import (
	"github.com/IcaroTARique/pr-locate-weather/internal/infra/api/cep"
	"github.com/IcaroTARique/pr-locate-weather/internal/infra/api/weather"
	"github.com/IcaroTARique/pr-locate-weather/internal/infra/webserver/handler"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	apiCep := cep.NewApiCep()
	apiWeather := weather.NewApiWeather()
	temperatureHandler := handler.NewApiTemperatureResponse(apiCep, apiWeather)

	r := chi.NewRouter()
	r.Get("/temperature/{cep}", temperatureHandler.GetTemperatureHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
