package models

type WeatherAPIResponse struct {
	Base   string `json:"base"`
	Clouds struct {
		All int64 `json:"all"`
	} `json:"clouds"`
	Cod   int64 `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   int64 `json:"dt"`
	ID   int64 `json:"id"`
	Main struct {
		FeelsLike float64 `json:"feels_like"`
		Humidity  int64   `json:"humidity"`
		Pressure  int64   `json:"pressure"`
		Temp      float64 `json:"temp"`
		TempMax   float64 `json:"temp_max"`
		TempMin   float64 `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
		ID      int64  `json:"id"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
		Type    int64  `json:"type"`
	} `json:"sys"`
	Timezone   int64 `json:"timezone"`
	Visibility int64 `json:"visibility"`
	Weather    []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
		ID          int64  `json:"id"`
		Main        string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   int64   `json:"deg"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}
