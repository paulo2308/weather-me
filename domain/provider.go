package domain

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Provider struct {
	Client *http.Client
}

func (p *Provider) GetCities(req Request) ([]byte, error) {
	url := fmt.Sprintf("https://revgeocode.search.hereapi.com/v1/revgeocode?limit=20&apiKey=xlHqo0XVNPT1tOb-1bTYAI-jJrvkpK-dHwJz9iMgXzo&in=circle:%s,%s;r=%d&types=city", req.Latitude, req.Longitude, req.Distancia)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p *Provider) GetPrevisao(cities string) ([]byte, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=EJ6UBL2JEQGYB3AA4ENASN62J&include=days,current&elements=datetime,tempmax,tempmin,temp,humidity,conditions,description,icon,precipprob,preciptype,snow,cloudcover&contentType=json&lang=pt", cities)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p *Provider) GetPrevisaoLite(cities string) ([]byte, error) {
	currentDate := time.Now().Truncate(24 * time.Hour)
	oneDayLater := currentDate.AddDate(0, 0, 1)
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s/?unitGroup=metric&key=EJ6UBL2JEQGYB3AA4ENASN62J&include=days,current&elements=datetime,tempmax,tempmin,temp,humidity,conditions,description,icon,precipprob,preciptype,snow,cloudcover&contentType=json&lang=pt", cities, currentDate.Format("2006-01-02"), oneDayLater.Format("2006-01-02"))
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
