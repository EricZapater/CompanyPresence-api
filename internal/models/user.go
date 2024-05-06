package models

type User struct {
	ID        string `json:"ID"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	Email     string `json:"Email"`
	IpAddress string `json:"IpAddress"`
	Active    bool   `json:"Active"`
}