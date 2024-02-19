package models

type Portfolio struct {
	UserId     string             `json:"userid"`
	Currencies map[string]float64 `json:"currencies"`
}
