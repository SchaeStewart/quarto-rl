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
	if x > 3 || y > 3 {
		return errors.New("cannot place piece outside of board bounds")
	}

	board := *b

	if board.IsWon() {
		return errors.New("cannot place a piece a won board")
	}

	if board[x][y] != nil {
		return errors.New("cannot place piece where a piece already exists")
	}

	board[x][y] = p

	return nil
}

func (bx *Board) IsWon() bool {
	b := *bx
	for i := 0; i < 4; i++ {
		// Col search
		if ShareAttributes(b[0][i], b[1][i], b[2][i], b[3][i]) {
			return true
		}
		// Row Search
		if ShareAttributes(b[i][0], b[i][1], b[i][2], b[i][3]) {
			return true
		}
	}

	// diaganal
	if ShareAttributes(b[0][0], b[1][1], b[2][2], b[3][3]) ||
		ShareAttributes(b[0][3], b[1][2], b[2][1], b[3][0]) {
		return true
	}

	return false
}

func (b Board) IsTied() bool {
	for i := range b {
		for _, x := range b[i] {
			if x == nil {
				return false
			}
		}
	}
	return true
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
