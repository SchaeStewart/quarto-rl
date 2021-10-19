package main

import "fmt"

type Piece struct {
	Tall   bool
	Black  bool
	Flat   bool
	Square bool
}

func (p *Piece) String() string {
	if p == nil {
		return "0000"
	}

	size := "T"
	if !p.Tall {
		size = "S"
	}

	color := "B"
	if !p.Black {
		color = "W"
	}

	top := "F"
	if !p.Flat {
		top = "D"
	}

	shape := "R" // R for rectangle
	if !p.Square {
		shape = "C"
	}

	return fmt.Sprintf("%s%s%s%s", size, color, top, shape)
}

func (p Piece) Print() {
	fmt.Print(p.String())
}

func IsIdentical(p, q Piece) bool {
	return p.Tall == q.Tall &&
		p.Black == q.Black &&
		p.Flat == q.Flat &&
		p.Square == q.Square
}

func ShareAttributes(p1, p2, p3, p4 *Piece) bool {
	if p1 == nil || p2 == nil || p3 == nil || p4 == nil {
		return false
	}
	if p1.Tall == p2.Tall && p1.Tall == p3.Tall && p1.Tall == p4.Tall {
		return true
	}
	if p1.Square == p2.Square && p1.Square == p3.Square && p1.Square == p4.Square {
		return true
	}
	if p1.Black == p2.Black && p1.Black == p3.Black && p1.Black == p4.Black {
		return true
	}
	if p1.Flat == p2.Flat && p1.Flat == p3.Flat && p1.Flat == p4.Flat {
		return true
	}
	return false
}

func CreatePieceSet() []Piece {
	return []Piece{
		{true, true, true, true},
		{true, false, true, true},
		{true, false, false, true},
		{true, false, false, false},
		{true, true, false, false},
		{true, true, true, false},
		{true, false, true, false},
		{true, true, false, true},
		{false, true, true, true},
		{false, false, true, true},
		{false, false, false, true},
		{false, false, false, false},
		{false, true, false, false},
		{false, true, true, false},
		{false, false, true, false},
		{false, true, false, true},
	}
}
