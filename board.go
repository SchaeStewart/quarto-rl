package main

type Board struct {
	Board [][]*Piece
}

func NewBoard () *Board {
	b := [][]*Piece{
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
	}
	return &Board{Board: b} 
} 