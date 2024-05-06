package models

type User struct {
	ID        string `json:"ID"`
	Name      string `json:"Name"`
	Surname   string `json:"Surname"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	IpAddress string `json:"IpAddress"`
	Active    bool   `json:"Active"`
}