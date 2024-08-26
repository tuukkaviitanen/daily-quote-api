package models

type Quote struct {
	Title  string `json:"title"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}
