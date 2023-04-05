package domain

import (
	"fmt"
	"net/http"
	"strconv"

)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var req request
	var err error
	req.city = queryParams.Get("cidade")
	req.initialDate = queryParams.Get("data_inicial")
	req.finaldate = queryParams.Get("data_fim")
	req.distance, err = strconv.Atoi(queryParams.Get("distancia"))
	if err != nil {
		return
	}
	req.temperature, err = strconv.Atoi(queryParams.Get("temperatura"))
	if err != nil {
		return
	}
	req.climate = queryParams.Get("clima")
	req.chanceOfRain, err = strconv.Atoi(queryParams.Get("chance_chuva"))
	if err != nil {
		return
	}
	req.typeOfTourism = queryParams.Get("tipo_turismo")
	resp, err := Get()
	fmt.Println(resp, err)

	w.WriteHeader(200)
	fmt.Fprint(w, resp)
}
