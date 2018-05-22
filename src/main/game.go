package main

import (
    "math/rand"
)

type Game struct {
    Players     []Player
}

func (g Game) AddPlayer(name string) Player {
    player := Player {
        numPlayers,
        name,
        rand.Float64() * 4000,
        rand.Float64() * 4000,
    }
    numPlayers++
    g.Players = append(g.Players, player)
    return player
}

func (g Game) Update() {

}
