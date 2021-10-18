package main

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	if len(b.Board) != 4 {
		t.Errorf("expected board length to be 4. instead got %d", len(b.Board))
	}

	for _, x := range b.Board {
		if len(x) != 4 {
			t.Errorf("expected row length to be 4. instead got %d", len(x))
		}
	}
}