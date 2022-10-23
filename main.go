package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askDraw() bool {
	fmt.Println("ドローしますか？ y/n")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	ans := input.Text()
	if ans == "y" {
		return true
	}
	return false
}

func setup(p, dl *player, dc *deck) {
	p.hit(*dc)
	dl.hit(*dc)
	fmt.Printf("%v,%v\n,%d\n", dl.name, dl.hand, dl.score())
	fmt.Printf("%v,%v\n,%d\n", p.name, p.hand, p.score())
}

//func pretty(p *player) {
//	fmt.Printf("name:%s\n", p.name)
//	for _, h := range p.hand {
//		fmt.Println("+-+")
//		fmt.Printf("|%v|\n", h.number)
//		fmt.Print("+-+\n\n")
//	}
//}

func hitOrCall(p, dl *player, dc *deck, hit bool) string {
	if hit {
		p.hit(*dc)
		if p.sumHand() > 21 {
			fmt.Printf("%v\n", p.sumHand())
			return "バーストしました"
		} else {
			fmt.Printf("%v,%v\n%v\n", dl.name, dl.hand, dl.score())
			fmt.Printf("%v,%v\n%v\n", p.name, p.hand, p.score())
			hitOrCall(p, dl, dc, askDraw())
		}
	}
	return strconv.Itoa(p.sumHand())
}

func main() {
	dc := newDeck()
	p := NewPlayer("player")
	dl := NewPlayer("dealer")
	setup(p, dl, dc)
	hitOrCall(p, dl, dc, askDraw())
	// result()
}
