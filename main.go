package main

import (
	"fmt"
)

func main() {
	fmt.Println("ヒットならy, コールならnを押してください")
	d := newDeck()
	p := NewPlayer()
	p.hit(*d)
	fmt.Printf("%v\n", p.hand)
	fmt.Printf("%v", d)
}
