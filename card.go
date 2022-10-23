package main

import "strconv"

type card struct {
	number int
	suit   int
}

func (c *card) toStringNum(int) string {
	switch c.number {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(c.number)
	}
}
