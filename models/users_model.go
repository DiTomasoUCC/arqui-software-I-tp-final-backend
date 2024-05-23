package models

type User struct {
	User_id   string `json:"Id"`
	Username  string `json:"Username"`
	Name      string `json:"Name"`
	Last_name string `json:"Last_name"`
	User_type string `json:"User_type"`
}
