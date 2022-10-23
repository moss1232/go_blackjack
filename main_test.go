package main

import (
	"fmt"
	"testing"
)

func TestPlayGame(t *testing.T) {
	dc := newDeck()
	p := NewPlayer("player")
	dl := NewPlayer("dealer")
	cases := []struct {
		name     string
		hit      bool
		expected string
	}{{
		name:     "success",
		hit:      true,
		expected: "true",
	},
	}

	for k, v := range cases {
		t.Run(fmt.Sprintf("#%d", k), func(t *testing.T) {
			got := hitOrCall(p, dl, dc, v.hit)
			if got != v.expected {
				t.Errorf("want %v, got %v", v.expected, got)
			}
		})
	}
}
