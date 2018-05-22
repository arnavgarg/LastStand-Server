package main

import (
	"net/http"
	"log"
	"fmt"

	"github.com/gorilla/mux"
	"encoding/json"
)

var game Game

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[ SUCCESS ] Request from ", r.RemoteAddr)
		decoder := json.NewDecoder(r.Body)
		var data AddPlayerData
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		// game.AddPlayer(data.name)

		json.NewEncoder(w).Encode(data)
	}).Methods("PUT");

	log.Fatal(http.ListenAndServe(":8080", router))

	go startGame()
}

func startGame() {

}

