package main

type AddPlayerData struct {
	Name 	string `json:"name"`
}

type ChangeData struct {
	Player	int `json:"player"`
	Log		[]Entry `json:"log"`
}

type Entry struct {
	Id		int `json:"id"`
	Extras	[]Extra `json:"extras"`
}

type Extra struct {
	Value	string `json:"value"`
}

type GameData struct {
	Players	[]Player `json:"players"`
}
