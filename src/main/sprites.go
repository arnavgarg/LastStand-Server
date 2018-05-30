package main

import "math/rand"

type Player struct {
	Id		int `json:"id"`
	Name    string `json:"name"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Health	int `json:"health"`
}

func (p *Player) moveRight() {
	if p.X < 3950 {
		p.X += 2;
	}
}

func (p *Player) moveLeft() {
	if p.X > 50 {
		p.X -= 2;
	}
}

func (p *Player) moveUp() {
	if p.Y > 50 {
		p.Y -= 2;
	}
}

func (p *Player) moveDown() {
	if p.Y < 3950 {
		p.Y += 2;
	}
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
