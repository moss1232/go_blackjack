package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askDraw() bool {
	fmt.Println("ヒット y/n")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println("----------------------------")
	ans := input.Text()
	if ans == "y" {
		return true
	} else if ans == "n" {
		return false
	} else {
		fmt.Print("コールします。")
	}
	return false
}

func setup(p, dl *player, dc *deck) {
	p.hit(*dc)
	dl.hit(*dc)
	p.prettyHand()
	dl.prettyHand()
}

func hitOrCall(p, dl *player, dc *deck, hit bool) string {
	if hit {
		p.hit(*dc)
		if p.sumHand() > 21 {
			fmt.Printf("%v\n", p.sumHand())
		} else {
			dl.prettyHand()
			p.prettyHand()
			hitOrCall(p, dl, dc, askDraw())
		}
	}
	return strconv.Itoa(p.sumHand())
}

func result(ps, dls int) {
	fmt.Print("RESULT: ")
	bj := 21
	if (ps > bj && dls > bj) || (ps > bj && dls <= bj) || (dls > ps) {
		fmt.Println("LOSE")
	} else if ps <= bj && ps == dls {
		fmt.Println("DRAW")
	} else {
		fmt.Println("WIN")
	}
}
