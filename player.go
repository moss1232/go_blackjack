package main

import "fmt"

type player struct {
	role string
	hand []card
}

func newPlayer(role string) *player {
	return &player{
		role: role,
		hand: []card{},
	}
}

func (p *player) renderHand() {
	fmt.Printf("%v ", p.role)
	fmt.Print("Card:")
	for _, h := range p.hand {
		switch h.suit {
		case 1:
			fmt.Print("♠")
		case 2:
			fmt.Print("♣")
		case 3:
			fmt.Print("♡")
		case 4:
			fmt.Print("♢")
		}
		fmt.Printf("%d", h.number)
	}
	fmt.Printf(" Score:%v\n", p.score())
	fmt.Println("----------------------------")
	return
}

func (p *player) hit(d deck) {
	p.hand = append(p.hand, d[0])
	d = append(d[:0], d[1:]...)
}

func (p *player) score() (sum int) {
	for _, v := range p.hand {
		sum += v.number
	}
	return
}
