package main

var numPlayers = 0

type Player struct {
	Id			int `json:"id"`
	Name        string `json:"name"`
	X           float32 `json:"x"`
	Y           float32 `json:"y"`
}
