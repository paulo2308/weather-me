package domain

type request struct {
	city          string
	forecastDays  int
}

type response struct {
	City          string
	TempMin       int
	TempMax       int
	Climate       string
	TypeOfTourism string
	Date          string
}
