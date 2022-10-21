package main

import (
	"math/rand"
	"time"
)

type deck []Card

var suit = [4]string{"Spade", "Clover", "Heart", "Diamond"}
var number = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func NewDeck() *deck {
	d := make(deck, 0, 52)
	for _, s := range suit {
		for _, n := range number {
			d = append(d, Card{Suit: s, Number: n})
		}
	}
	return &d
}

func (d deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
