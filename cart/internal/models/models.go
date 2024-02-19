package models

type Cart struct {
	UserId   string         `json:"userid"`
	Products map[string]int `json:"products"`
}
