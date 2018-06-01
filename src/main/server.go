package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
)

var game Game

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data AddPlayerData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		json.NewEncoder(w).Encode(AddPlayerResponse {
			game.AddPlayer(data.Name),
			game.Rocks,
		})
		fmt.Println("[ SUCCESS ] New Player Joined from ", r.RemoteAddr)
	}).Methods("PUT");

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data ChangeData
		err := json.NewDecoder(r.Body).Decode(&data)
		if (err != nil) {
			panic(nil)
		}
		defer r.Body.Close()

		game.ApplyChanges(data)
		json.NewEncoder(w).Encode(game.GetGameData())
	}).Methods("POST")

	startGame()

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("[ SUCCESS ] Listening on Port 8080")
}

func startGame() {
	game = Game {
		make([]Player, 0),
		generateRocks(),
	}
}

