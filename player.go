package main

type player struct {
	hand []card
}

func NewPlayer() *player {
	return &player{
		hand: []card{},
	}
}

func (p *player) stand() {

}

func (p *player) hit(d deck) {
	p.hand = append(p.hand, d[0])
}
