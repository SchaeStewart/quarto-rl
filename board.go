package main

import (
	"errors"
	"fmt"
)

type Board [][]*Piece

func NewBoard() Board {
	var b Board = [][]*Piece{
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
	}
	return b
}

func (b *Board) SetPiece(x, y int, p *Piece) error {
	if x > 4 || y > 4 {
		return errors.New("cannot place piece outside of board bounds")
	}

	board := *b

	if board[x][y] != nil {
		return errors.New("cannot place piece where a piece already exists")
	}

	board[x][y] = p

	return nil
}

func (b Board) String() string {
	result := ""
	for _, x := range b {
		result += "|"
		for _, y := range x {
			result += fmt.Sprintf("%s|", y)
		}
		result += "\n"
	}
	return result
}

func (b Board) PrintBoard() {
	fmt.Print(b)
}
