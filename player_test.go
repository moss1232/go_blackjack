package main

import "testing"

func TestScore(t *testing.T) {
	p := NewPlayer("player")
	got := p.score()
	if got > 21 {
		t.Fatal("fail")
	}
}
