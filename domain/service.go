package domain

import "module/provider"

func Get() (provider.Response, error) {
	resp, err := provider.GetDataWeather()
	if err != nil {
		return provider.Response{}, err
	}
	return resp, nil
}