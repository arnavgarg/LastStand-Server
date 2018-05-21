package main

import (
	"net/http"
	"log"
	"fmt"

	"github.com/gorilla/mux"
	"encoding/json"
)

func main() {
	gameChannel := make(chan Game)

	router := mux.NewRouter()
	router.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		game := <-gameChannel
		decoder := json.NewDecoder(r.Body)
		fmt.Println("Connection from ", r.RemoteAddr)

		var data AddPlayerData
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		// game.AddPlayer(data.name)
		gameChannel <-game
		json.NewEncoder(w).Encode(data)
		fmt.Println(data.name)
	}).Methods("PUT");

	log.Fatal(http.ListenAndServe(":8080", router))

	go startGame(gameChannel)
}

func startGame(gameChannel chan Game) {
	gameChannel <- Game {

	}
}

