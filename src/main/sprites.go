package main

import (
	"math/rand"
	"math"
)

type Player struct {
	Id		int `json:"id"`
	Name    string `json:"name"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Angle	float64	`json:"angle"`
	Health	int `json:"health"`
	Status	int `json:"status"`
}

func (p *Player) moveRight() {
	if p.X < 3950 {
		p.X += 5;
	}
	game.CheckCollisions(p)
}

func (p *Player) moveLeft() {
	if p.X > 50 {
		p.X -= 5;
	}
	game.CheckCollisions(p)
}

func (p *Player) moveUp() {
	if p.Y > 50 {
		p.Y -= 5;
	}
	game.CheckCollisions(p)
}

func (p *Player) moveDown() {
	if p.Y < 3950 {
		p.Y += 5;
	}
	game.CheckCollisions(p)
}

func (p *Player) collisionRock(r Rock) bool {
	return math.Sqrt(math.Pow(p.X - r.X, 2) + math.Pow(p.Y - r.Y, 2)) <= 90
}

type Rock struct {
	X		float64 `json:"x"`
	Y		float64 `json:"y"`
}

func generateRocks() []Rock {
	rocks := make([]Rock, 250)
	for i := 0; i < len(rocks); i++ {
		rocks[i] = Rock {
			rand.Float64() * 3950 + 25,
			rand.Float64() * 3950 + 25,
		}
	}
	return rocks
}
