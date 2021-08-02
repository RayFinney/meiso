package models

import "time"

type Plant struct {
	Name        string           `json:"name"`
	BotanicName *string          `json:"botanicName"`
	Description *string          `json:"description"`
	Family      *string          `json:"family"`
	Origin      *string          `json:"origin"`
	MaxHeight   *int64           `json:"maxHeight"`
	WinterHard  bool             `json:"winterHard"`
	Extras      *string          `json:"extras"`
	CarePlan    PlantCare        `json:"carePlan"`
	CareTasks   []PlantCareTasks `json:"careTasks"`
}

type PlantCare struct {
	Lux            float64 `json:"lux"`    // Lux per day (Lux * time.Duration)
	MaxLux         float64 `json:"maxLux"` // Peak Lux
	MinTemperature float64 `json:"minTemperature"`
	MaxTemperature float64 `json:"maxTemperature"`
	MinHumidity    float64 `json:"minHumidity"`
	MaxHumidity    float64 `json:"maxHumidity"`
}

type PlantCareTasks struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Interval    time.Duration `json:"interval"`
}
