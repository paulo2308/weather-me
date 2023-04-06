package app

import (
	"net/http"

	"module/domain"
)

func Routes() {
	http.HandleFunc("/previsoes", domain.GetWeather)
	http.ListenAndServe(":8080", nil)
}
