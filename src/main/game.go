package main

import (
    "math/rand"
)

type Game struct {
    Players     []Player
    Rocks       []Rock
}

func (g *Game) AddPlayer(name string) Player {
    player := Player {
        len(g.Players),
        name,
        rand.Float64() * 3000 + 500,
        rand.Float64() * 4000 + 500,
    }
    g.Players = append(g.Players, player)
    return player
}

func (g *Game) ApplyChanges(changes ChangeData) {
    for _,entry := range changes.Log {
        switch entry.Id {
        case 0:
            game.Players[changes.Player].moveUp()
        case 1:
            game.Players[changes.Player].moveRight()
        case 2:
            game.Players[changes.Player].moveDown()
        case 3:
            game.Players[changes.Player].moveLeft()
        }
    }
}

func (g *Game) GetGameData() GameData {
    return GameData {
        g.Players,
        g.Rocks,
    }
}

