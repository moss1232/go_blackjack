package main

func main() {
	dc := newDeck()
	p := newPlayer("player")
	dl := newPlayer("dealer")
	setup(p, dl, dc)
	hitOrCall(p, dl, dc, askHit())
	result(p, dl, dc)
}
