package main

type AddPlayerData struct {
	Name 	string `json:"name"`
}

type AddPlayerResponse struct {
	Player	Player `json:"player"`
	Rocks	[]Rock `json:"rocks"`
}

type ChangeData struct {
	Player	int `json:"player"`
	Log		[]Entry `json:"log"`
}

type Entry struct {
	Id		int `json:"id"`
	Extras	[]string `json:"extras"`
}

type GameData struct {
	Players	[]Player `json:"players"`
}
