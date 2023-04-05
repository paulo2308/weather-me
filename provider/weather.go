package provider

func GetDataWeather() (Response, error) {
	s := []string{"sbc","itapevi","interlagos"}
	r := Response{
		Cities: s,
	}

	return r, nil
}