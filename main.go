package main

import (
	"fmt"
)

func msgHitOrCall() {
	fmt.Println("ヒットならy, コールならnを押してください")
}

func setup(p, dl *player, dc *deck) {
	p.hit(*dc)
	dl.hit(*dc)
	pretty(p)
	pretty(dl)
}

func pretty(p *player) {
	fmt.Printf("name:%s\n", p.name)
	for _, h := range p.hand {
		fmt.Println("+-+")
		fmt.Printf("|%s|\n", h.number)
		fmt.Print("+-+\n\n")
	}
}

func main() {
	dc := newDeck()
	fmt.Printf("%v", dc)
	p := NewPlayer("player")
	dl := NewPlayer("dealer")
	setup(p, dl, dc)
	//msgHitOrCall()
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//it := input.Text()
	//if it == "Y" {
	//	p.hit(*dc)
	//}
	//if it == "n" {
	//
	//}
	//fmt.Printf("%+v", input.Text())
	//p.hit(*d)
	//fmt.Printf("%v\n", p.hand)
	//fmt.Printf("%v", d)
}
