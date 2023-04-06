package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var req request
	var err error
	req.city = queryParams.Get("cidade")
	req.forecastDays, err = strconv.Atoi(queryParams.Get("dias_previsao"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	resp, err := Get()
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
