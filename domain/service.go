package domain

func Get() ([]response, error) {
	var resp []response
	resp = []response{
		{
			City:          "Sao Sebastiao",
			TempMin:       24,
			TempMax:       30,
			Climate:       "ensolarado",
			TypeOfTourism: "praia",
			Date:          "06/05/2023",
		},
	}
	return resp, nil
}
