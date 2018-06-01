package main

import (
    "math/rand"
    "strconv"
    "math"
)

type Game struct {
    Players     []Player
    Rocks       []Rock
}

func (g *Game) AddPlayer(name string) Player {
    player := Player {
        len(game.Players),
        name,
        rand.Float64() * 3000 + 500,
        rand.Float64() * 3000 + 500,
        1.57,
        100,
        1,
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
        case 4:
            angle,_ := strconv.ParseFloat(entry.Extras[0], 64)
            game.Shoot(game.Players[changes.Player], angle)
        case 5:
            game.Players[changes.Player].Angle,_ = strconv.ParseFloat(entry.Extras[0], 64)
        }
    }
}

func (g *Game) CheckCollisions(p *Player) {
    for _,r := range g.Rocks {
        if p.collisionRock(r) {
            p.Status = 0;
        }
    }
}

func (g *Game) Shoot(player Player, angle float64) {
    for _,p := range g.Players {
        if player.Id == p.Id {
            continue
        }

        distance := int(math.Sqrt(math.Pow(p.X - player.X, 2) + math.Pow(p.Y - player.Y, 2)))
        if distance < 500 {
            a := math.Atan((player.Y - p.Y)/(player.X - p.X))
            if angle < a + 0.5 && angle > a - 0.5 {
                p.Health -= int(20 * (1 - (float64(distance)/100) + .3))
                if p.Health <= 0 {
                    p.Status = 0
                }
            }
        }
    }
}

func (g *Game) GetGameData() GameData {
    return GameData {
        g.Players,
    }
}



