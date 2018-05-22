package main

import (
	"strconv"
)

type AddPlayerData struct {
	Name string `json:"name"`
}

type AddPlayerResponse struct {
	Id			string `json:"id"`
	Name        string `json:"name"`
	X           string `json:"x"`
	Y           string `json:"y"`
}

func getAddPlayerResponse(data Player) AddPlayerResponse {
	return AddPlayerResponse {
		strconv.FormatInt(int64(data.Id), 10),
		data.Name,
		strconv.FormatFloat(data.X, 'f', 2, 64),
		strconv.FormatFloat(data.Y, 'f', 2, 64),
	}
}