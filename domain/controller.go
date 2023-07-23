package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var req Request
	var err error
	req.Latitude = queryParams.Get("latitude")
	req.Longitude = queryParams.Get("longitude")
	req.Cidade = queryParams.Get("cidade")

	dist := queryParams.Get("distancia")
	if dist != "" {
		req.Distancia, err = strconv.Atoi(dist)
		if err != nil {
			w.WriteHeader(400)
			return
		}
	}

	clima := queryParams.Get("clima")
	switch clima {
	case "ceu_limpo", "sol_entre_nuvens", "nublado":
		req.Clima = clima
	default:
		req.Clima = ""
	}

	req.Ordenacao = queryParams.Get("sort")

	req.Pagina, err = strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	req.ItensPorPagina, err = strconv.Atoi(queryParams.Get("itens_per_page"))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	resp, err := GetCitiesService(&req)
	if err != nil {
		return
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, string(respJson))
}
