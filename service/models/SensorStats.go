package models

type SensorStats struct {
	Temp     float64 `json:"temp"`
	FTemp    float64 `json:"fTemp"`
	Humidity float64 `json:"humidity"`
	Lux      float64 `json:"lux"`
}
