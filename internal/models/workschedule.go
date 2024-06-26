package models

import "time"

type WorkSchedule struct {
	ID                 string    `json:"ID"`
	UserID             string    `json:"UserID"`
	NormalWorkingHours float32   `json:"NormalWorkingHours"`
	NormalStartTime    time.Time `json:"NormalStartTime"`
	NormalNoonRest float32 `json:"NormalNoonRest"`
	FridayWorkingHours float32   `json:"FridayWorkingHours"`
	FridayStartTime    time.Time `json:"FridayStartTime"`
}