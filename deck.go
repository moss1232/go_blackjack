package main

import (
	"math/rand"
	"time"
)

type deck []card

type card struct {
	number int
	suit   int
}

//var suit = [4]string{"Spade", "Clover", "Heart", "Diamond"}
//var number = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func newDeck() *deck {
	d := make(deck, 0, 52)
	for s := 1; s <= 4; s++ {
		for n := 1; n <= 13; n++ {
			d = append(d, card{suit: s, number: n})
		}
	}
	d.shuffle()
	return &d
}

func (d *deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}
