package model

type TouristProgress struct {
	Id     string `json:"Id"`
	UserId string `json:"Userid"`
	Xp     string `json:"Xp"`
	Level  string `json:"Level"`
}

type TouristProgressDto struct {
	Xp    int
	Level int
}
