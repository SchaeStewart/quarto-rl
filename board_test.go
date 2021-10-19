package main

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	if len(b) != 4 {
		t.Errorf("expected board length to be 4. instead got %d", len(b))
	}

	for _, x := range b {
		if len(x) != 4 {
			t.Errorf("expected row length to be 4. instead got %d", len(x))
		}
	}
}

func TestSetPiece(t *testing.T) {
	b := NewBoard()
	p := Piece{true, true, true, true}
	b.SetPiece(0, 0, &p)

	if (!IsIdentical(*b[0][0], Piece{true, true, true, true})) {
		t.Errorf("failed to update board\n%v, Piece: %v", b, p.String())
	}

	err := b.SetPiece(4, 4, &Piece{true, true, true, true})
	if err == nil {
		t.Error("expected err when setting piece outside of board bounds")
	}
}

func TestBoardString(t *testing.T) {
	b := NewBoard()

	emptyBoardString := "|0000|0000|0000|0000|\n|0000|0000|0000|0000|\n|0000|0000|0000|0000|\n|0000|0000|0000|0000|\n"

	if b.String() != emptyBoardString {
		t.Errorf("expected empty board instead got %s", b.String())
	}

	nonEmptyBoardString := "|TBFR|0000|0000|0000|\n|0000|0000|0000|0000|\n|0000|0000|0000|0000|\n|0000|0000|0000|TBFC|\n"
	b.SetPiece(0, 0, &Piece{true, true, true, true})
	b.SetPiece(3, 3, &Piece{true, true, true, false})
	if b.String() != nonEmptyBoardString {
		t.Errorf("expected board to be:\n%s \ninstead got:\n%s", nonEmptyBoardString, b.String())
	}
}

func TestBoardWon(t *testing.T) {
	// Row check
	b := NewBoard()
	b.SetPiece(0, 0, &Piece{true, false, false, false})
	b.SetPiece(0, 1, &Piece{true, true, false, false})
	b.SetPiece(0, 2, &Piece{true, false, true, false})
	b.SetPiece(0, 3, &Piece{true, false, false, true})

	if !b.IsWon() {
		t.Errorf("expected board to be in a win state\n%v", b)
	}

	// Col check
	b = NewBoard()
	b.SetPiece(0, 0, &Piece{true, false, false, false})
	b.SetPiece(1, 0, &Piece{true, true, false, false})
	b.SetPiece(2, 0, &Piece{true, false, true, false})
	b.SetPiece(3, 0, &Piece{true, false, false, true})

	if !b.IsWon() {
		t.Errorf("expected board to be in a win state\n%v", b)
	}

	// Diaganal check
	b = NewBoard()
	b.SetPiece(0, 0, &Piece{true, false, false, false})
	b.SetPiece(1, 1, &Piece{true, true, false, false})
	b.SetPiece(2, 2, &Piece{true, false, true, false})
	b.SetPiece(3, 3, &Piece{true, false, false, true})

	if !b.IsWon() {
		t.Errorf("expected board to be in a win state\n%v", b)
	}
}
