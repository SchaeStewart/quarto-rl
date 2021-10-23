package main

import (
	"math/rand"
)

type Agent interface {
	SelectPiece() Piece
	PlayPiece() (x, y int)
}

type RandomAgent struct {
	Environment *Game
	Seed int64
}

func (a RandomAgent) SelectPiece() Piece {
	rand.Seed(a.Seed)
	pieceIdx := rand.Intn(len(a.Environment.Pieces))
	return a.Environment.Pieces[pieceIdx]
}

func (a RandomAgent) PlayPiece() (x, y int) {
	rand.Seed(a.Seed)

	// Find the available spaces
	openSpaces := make([]struct{X, Y int}, 0)
	for i := range a.Environment.Board {
		for j := range a.Environment.Board[i] {
			if a.Environment.Board[i][j] == nil {
				openSpaces = append(openSpaces, struct{X int; Y int}{i,j})
			}
		}
	}

	spaceIdx := rand.Intn(len(openSpaces)) 
	space := openSpaces[spaceIdx]

	return space.X, space.Y
}

func NewRandomAgent(seed int64, environment *Game) RandomAgent {
	return RandomAgent{
		Seed: seed,
		Environment: environment,
	}
}