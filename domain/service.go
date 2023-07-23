package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

func GetCitiesService(req *Request) ([]WeatherResponse, error) {
	if req.Distancia == 0 {
		req.Distancia = 1
	}

	cities, err := ListOfCity(*req)
	if err != nil {
		return nil, err
	}

	pag := paginacao(req, cities)
	sort, err := GetPrevisoes(pag)
	if err != nil {
		return nil, err
	}

	return Ordenacao(req, sort), nil
}

func ListOfCity(req Request) (ListCities, error) {
	var cities ListCities
	provider := Provider{
		Client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
	body, err := provider.GetCities(req)
	if err != nil {
		return ListCities{}, err
	}

	err = json.Unmarshal(body, &cities)
	if err != nil {
		return ListCities{}, err
	}

	if err != nil {
		return ListCities{}, err
	}

	return cities, nil
}

func GetPrevisoes(pag []Item) ([]WeatherResponse, error) {

	var prev WeatherResponse
	var allPrev []WeatherResponse
	provider := Provider{
		Client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}

	for _, i := range pag {
		var body []byte
		var err error

		if len(pag) > 1 {
			fmt.Println(i.Address.City)
			body, err = provider.GetPrevisaoLite(i.Address.City)
			if err != nil {
				return []WeatherResponse{}, err
			}
		} else {
			fmt.Println(i.Address.City)
			body, err = provider.GetPrevisao(i.Address.City)
			if err != nil {
				return []WeatherResponse{}, err
			}
		}

		json.Unmarshal(body, &prev)
		// if err != nil {
		// 	fmt.Println("aqui")
		// 	return []WeatherResponse{}, err
		// }

		prev.Distance = i.Distance
		allPrev = append(allPrev, prev)
	}

	return allPrev, nil
}

func paginacao(r *Request, lc ListCities) []Item {

	totalItems := len(lc.Items)
	itemsPerPage := r.ItensPorPagina
	currentPage := r.Pagina
	startIndex := (currentPage - 1) * itemsPerPage
	endIndex := startIndex + itemsPerPage

	if endIndex > totalItems {
		endIndex = totalItems
	}

	pagedItems := lc.Items[startIndex:endIndex]

	return pagedItems
}

func Ordenacao(r *Request, prev []WeatherResponse) []WeatherResponse {
	switch r.Ordenacao {
	case "distancia:asc":
		sort.Sort(byDist(prev))
		fmt.Println(prev)
		fmt.Println("================================================")
	case "distancia:desc":
		sort.Sort(sort.Reverse(byDist(prev)))
	case "temperatura:asc":
		for _, p := range prev {
			sort.Sort(byTemp(p.Days))
		}
	case "temperatura:desc":
		for _, p := range prev {
			sort.Sort(sort.Reverse(byTemp(p.Days)))
		}
		fmt.Println(prev)
		fmt.Println("================================================")
	}

	return prev
}

type byTemp []Day

func (bt byTemp) Len() int           { return len(bt) }
func (bt byTemp) Swap(i, j int)      { bt[i], bt[j] = bt[j], bt[i] }
func (bt byTemp) Less(i, j int) bool { return bt[i].Temp < bt[j].Temp }

type byDist []WeatherResponse

func (bd byDist) Len() int           { return len(bd) }
func (bd byDist) Swap(i, j int)      { bd[i], bd[j] = bd[j], bd[i] }
func (bd byDist) Less(i, j int) bool { return bd[i].Distance < bd[j].Distance }
