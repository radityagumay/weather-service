package models

type Weather struct {
	Temperature string `json:"temperature"`
	Condition   string `json:"condition"`
	Location    string `json:"location"`
}
