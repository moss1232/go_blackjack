package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("ヒットならy, コールならnを押してください")
	//d := newDeck()
	//p := NewPlayer()
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Printf("%+v", input.Text())
	//p.hit(*d)
	//fmt.Printf("%v\n", p.hand)
	//fmt.Printf("%v", d)
}
