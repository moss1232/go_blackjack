package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askHit() bool {
	fmt.Println("hit? y/n")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println("----------------------------")
	ans := input.Text()
	if ans == "y" {
		return true
	}
	fmt.Println("call")
	return false
}

func setup(p, dl *player, dc deck) {
	dl.hit(dc)
	p.hit(dc)
	p.renderHand()
	dl.renderHand()
}

func hitOrCall(p, dl *player, dc deck, hit bool) string {
	if hit {
		p.hit(dc)
		if p.score() > 21 {
			fmt.Printf("%v\n", p.score())
		} else {
			dl.renderHand()
			p.renderHand()
			hitOrCall(p, dl, dc, askHit())
		}
	}
	return strconv.Itoa(p.score())
}

func result(p, dl *player, dc deck) {
	for dls := 0; dls < 17; {
		dl.hit(dc)
		dls = dl.score()
	}
	p.renderHand()
	dl.renderHand()
	fmt.Print("RESULT: ")
	bj := 21
	if (p.score() > bj && dl.score() > bj) || (p.score() > bj && dl.score() <= bj) || (p.score() < dl.score()) {
		fmt.Println("YOU LOSE")
	} else if p.score() <= bj && p.score() == dl.score() {
		fmt.Println("DRAW")
	} else {
		fmt.Println("YOU WIN!")
	}
}
