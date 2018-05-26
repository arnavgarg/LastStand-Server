package main

import "math/rand"

type Player struct {
	Id		int `json:"id"`
	Name    string `json:"name"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
}

func (p *Player) moveRight() {
	p.X += 2;
}

func (p *Player) moveLeft() {
	p.X -= 2;
}

func (p *Player) moveUp() {
	p.Y -= 2;
}

func (p *Player) moveDown() {
	p.Y += 2;
}

type Rock struct {
	X		float64 `json:x`
	Y		float64 `json:y`
}

func generateRocks() []Rock {
	rocks := make([]Rock, 1000)
	for i := 0; i < len(rocks); i++ {
		rocks[i] = Rock {
			rand.Float64() * 3980 + 10,
			rand.Float64() * 3980 + 10,
		}
	}
	return rocks
}
