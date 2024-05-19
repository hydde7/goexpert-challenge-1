package models

type FreeWeatherPayload struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Current struct {
	LastUpdatedEpoch int     `json:"last_updated_epoch"`
	LastUpdated      string  `json:"last_updated"`
	TempC            float64 `json:"temp_c"`
	TempF            float64 `json:"temp_f"`
	IsDay            int     `json:"is_day"`
	Condition        struct {
		Text string `json:"text"`
		Icon string `json:"icon"`
		Code int    `json:"code"`
	} `json:"condition"`
	WindMph    float64 `json:"wind_mph"`
	WindKph    float64 `json:"wind_kph"`
	WindDegree int     `json:"wind_degree"`
	WindDir    string  `json:"wind_dir"`
	PressureMb float64 `json:"pressure_mb"`
	PressureIn float64 `json:"pressure_in"`
	PrecipMm   float64 `json:"precip_mm"`
	PrecipIn   float64 `json:"precip_in"`
	Humidity   int     `json:"humidity"`
	Cloud      int     `json:"cloud"`
	FeelslikeC float64 `json:"feelslike_c"`
	FeelslikeF float64 `json:"feelslike_f"`
	VisKm      float64 `json:"vis_km"`
	VisMiles   float64 `json:"vis_miles"`
	Uv         float64 `json:"uv"`
	GustMph    float64 `json:"gust_mph"`
	GustKph    float64 `json:"gust_kph"`
	AirQuality struct {
		Co           float64 `json:"co"`
		No2          float64 `json:"no2"`
		O3           float64 `json:"o3"`
		So2          float64 `json:"so2"`
		Pm25         float64 `json:"pm2_5"`
		Pm10         float64 `json:"pm10"`
		UsEpaIndex   float64 `json:"us-epa-index"`
		GbDefraIndex float64 `json:"gb-defra-index"`
	} `json:"air_quality"`
}
