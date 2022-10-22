package main

type player struct {
	name string
	hand []card
}

func NewPlayer(name string) *player {
	return &player{
		name: name,
		hand: []card{},
	}
}

func (p *player) stand() {

}

func (p *player) hit(d deck) {
	p.hand = append(p.hand, d[0])
	d = append(d[:0], d[1:]...)
}
