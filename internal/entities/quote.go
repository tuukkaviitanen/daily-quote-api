package entities

type Quote struct {
	Id     int `gorm:"primary_key"`
	Quote  string
	Author string
}
