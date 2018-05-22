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
		decoder := json.NewDecoder(r.Body)
		var data AddPlayerData
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		json.NewEncoder(w).Encode(getAddPlayerResponse(game.AddPlayer(data.Name)))
		fmt.Println("[ SUCCESS ] New Player Joined from ", r.RemoteAddr)
	}).Methods("PUT");

	log.Fatal(http.ListenAndServe(":8080", router))

	go startGame()
}

func startGame() {

}

