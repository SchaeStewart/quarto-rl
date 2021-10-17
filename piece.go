package main
type Piece struct {
	Tall   bool
	Black  bool
	Flat   bool
	Square bool
}

func (p Piece) IsIdentical(q Piece) bool {
	return p.Tall == q.Tall &&
		p.Black == q.Black &&
		p.Flat == q.Flat &&
		p.Square == q.Square
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