package main

func main() {
	dc := newDeck()
	p := NewPlayer("player")
	dl := NewPlayer("dealer")
	setup(p, dl, dc)
	hitOrCall(p, dl, dc, askDraw())
	result(p.score(), dl.score())
}
