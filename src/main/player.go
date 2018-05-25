package main

type Player struct {
	Id			int `json:"id"`
	Name        string `json:"name"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
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
