package models

import "time"

type TimeTracking struct {
	ID          string `json:"ID"`
	UserID      string `json:"UserID"`
	WorkingDate time.Time `json:"WorkingDate"`
	ClockIn time.Time `json:"ClockIn"`
	ClockOut time.Time `json:"ClockOut"`
	IpAddress string `json:"IpAddress"`
}