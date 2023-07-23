package domain

type Request struct {
	Latitude       string
	Longitude      string
	Cidade         string
	Distancia      int
	Clima          string
	Ordenacao      string
	Pagina         int
	ItensPorPagina int
}

type Response struct {
	City          string
	TempMin       int
	TempMax       int
	Climate       string
	TypeOfTourism string
	Date          string
}

type Item struct {
	Title        string   `json:"title"`
	ID           string   `json:"id"`
	ResultType   string   `json:"resultType"`
	LocalityType string   `json:"localityType"`
	Address      Address  `json:"address"`
	Position     Position `json:"position"`
	Distance     int      `json:"distance"`
}

type Address struct {
	Label       string `json:"label"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	StateCode   string `json:"stateCode"`
	State       string `json:"state"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ListCities struct {
	Items []Item `json:"items"`
}

// ===================================================
type Hour struct {
	Datetime string `json:"datetime"`
	// DatetimeEpoch  int     `json:"datetimeEpoch"`
	Temp float64 `json:"temp"`
	// Feelslike      float64 `json:"feelslike"`
	Humidity float64 `json:"humidity"`
	// Dew            float64 `json:"dew"`
	// Precip         float64 `json:"precip"`
	Precipprob float64 `json:"precipprob"`
	Snow       float64 `json:"snow"`
	// Snowdepth      float64 `json:"snowdepth"`
	// Windgust       float64 `json:"windgust"`
	// Windspeed      float64 `json:"windspeed"`
	// Winddir        float64 `json:"winddir"`
	// Pressure       float64 `json:"pressure"`
	Visibility float64 `json:"visibility"`
	Cloudcover float64 `json:"cloudcover"`
	// Solarradiation float64 `json:"solarradiation"`
	// Solarenergy    float64 `json:"solarenergy"`
	// Uvindex        float64 `json:"uvindex"`
	// Severerisk     float64 `json:"severerisk"`
	Conditions string `json:"conditions"`
	Icon       string `json:"icon"`
	// Stations       []string `json:"stations"`
	// Source         string  `json:"source"`
}

type Day struct {
	Datetime string `json:"datetime"`
	// DatetimeEpoch  int     `json:"datetimeEpoch"`
	Tempmax float64 `json:"tempmax"`
	Tempmin float64 `json:"tempmin"`
	Temp    float64 `json:"temp"`
	// Feelslikemax   float64 `json:"feelslikemax"`
	// Feelslikemin   float64 `json:"feelslikemin"`
	// Feelslike      float64 `json:"feelslike"`
	// Dew            float64 `json:"dew"`
	Humidity float64 `json:"humidity"`
	// Precip         float64 `json:"precip"`
	Precipprob float64 `json:"precipprob"`
	// PrecipType 	   float64 `json:"preciptype"`
	// Precipcover    float64 `json:"precipcover"`
	Snow float64 `json:"snow"`
	// Snowdepth      float64 `json:"snowdepth"`
	// Windgust       float64 `json:"windgust"`
	// Windspeed      float64 `json:"windspeed"`
	// Winddir        float64 `json:"winddir"`
	// Pressure       float64 `json:"pressure"`
	Cloudcover float64 `json:"cloudcover"`
	// Visibility     float64 `json:"visibility"`
	// Solarradiation float64 `json:"solarradiation"`
	// Solarenergy    float64 `json:"solarenergy"`
	// Uvindex        float64 `json:"uvindex"`
	// Severerisk     float64 `json:"severerisk"`
	// Sunrise        string  `json:"sunrise"`
	// SunriseEpoch   int     `json:"sunriseEpoch"`
	// Sunset         string  `json:"sunset"`
	// SunsetEpoch    int     `json:"sunsetEpoch"`
	// Moonphase      float64 `json:"moonphase"`
	Conditions  string `json:"conditions"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	// Stations       []string `json:"stations"`
	// Source         string  `json:"source"`
	//Hours          []Hour  `json:"hours"`
}

type WeatherResponse struct {
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Distance        int     `json:"distance"`
	Timezone        string  `json:"timezone"`
	Tzoffset        float64 `json:"tzoffset"`
	Description     string  `json:"description"`
	Days            []Day   `json:"days"`
}
