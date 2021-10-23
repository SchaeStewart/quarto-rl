package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type GameMode int

const (
	NotStarted GameMode = iota
	Playing
	Complete
)

type Player int

const (
	Player1 Player = iota
	Player2
)

type Game struct {
	Board        Board
	Pieces       []Piece
	PlayedPieces []Piece
	GameMode     GameMode
	// The piece about to be played
	CurrentPiece  *Piece
	CurrentPlayer Player
}

func NewGame() *Game {
	return &Game{
		Board:        NewBoard(),
		Pieces:       CreatePieceSet(),
		PlayedPieces: []Piece{},
		GameMode:     NotStarted,
		CurrentPiece: nil,
	}
}

func (g *Game) PlayPiece(x, y int) error {
	err := g.Board.SetPiece(x, y, g.CurrentPiece)
	if err != nil {
		return err
	}
	g.CurrentPiece = nil
	return nil
}

func (g *Game) SelectPiece(p Piece) error {
	for _, q := range g.PlayedPieces {
		if IsIdentical(p, q) {
			return errors.New("piece has already been played")
		}
	}

	g.CurrentPiece = &p
	g.PlayedPieces = append(g.PlayedPieces, p)

	return nil
}

func (g *Game) PrintBoard() {
	g.Board.PrintBoard()
}

func (g *Game) IncrementPlayer() {
	if g.CurrentPlayer == Player1 {
		g.CurrentPlayer = Player2
		return
	}
	g.CurrentPlayer = Player1
}

func (g *Game) Play(x, y int, nextPiece Piece) error {
	if err := g.PlayPiece(x, y); err != nil {
		return err
	}
	if err := g.SelectPiece(nextPiece); err != nil {
		return err
	}

	if g.Board.IsWon() {
		return nil
	}

	g.IncrementPlayer()
	return nil
}

func GetPieceFromInput(reader *bufio.Reader) *Piece {
	fmt.Print("Enter your piece (T|S, B|W, F|D, R|C): ")
	pieceStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		GetPieceFromInput(reader)
	}
	p, err := StringToPiece(pieceStr[:4])
	if err != nil {
		fmt.Println(err.Error())
		GetPieceFromInput(reader)
	}
	return p
}

func GetLocationFromInput(reader *bufio.Reader) (x, y int) {
	fmt.Print("Enter your X value: ")
	xStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return GetLocationFromInput(reader)
	}
	fmt.Print("Enter your Y value: ")
	yStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return GetLocationFromInput(reader)
	}
	x64, err := strconv.ParseInt(string(xStr[0]), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return GetLocationFromInput(reader)
	}
	y64, err := strconv.ParseInt(string(yStr[0]), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return GetLocationFromInput(reader)
	}

	if x > 3 || y > 3 || x < 0 || y < 0 {
		fmt.Println("value must be 0 - 3")
		return GetLocationFromInput(reader)
	}

	return int(x64), int(y64)
}

func (g *Game) Loop(rd io.Reader) {
	reader := bufio.NewReader(rd)
	// TODO: put setting initial piece into function
	piece := GetPieceFromInput(reader)
	if piece == nil {
		fmt.Println("exiting game")
		return
	}
	g.CurrentPiece = piece
	g.IncrementPlayer()

	for !g.Board.IsWon() {
		// Start turn
		fmt.Println("Current player:", g.CurrentPlayer)
		fmt.Println("Current piece:", g.CurrentPiece)
		g.Board.PrintBoard()

		// Play piece
		x, y := GetLocationFromInput(reader)
		if err := g.PlayPiece(x, y); err != nil {
			fmt.Println(err.Error())
			continue
		}
		g.Board.PrintBoard()

		if g.Board.IsWon() {
			fmt.Println("Winning player:", g.CurrentPlayer)
			continue
		}

		if g.Board.IsTied() {
			fmt.Println("Game tied")
			continue
		}

		// Select next piece
		nextPiece := GetPieceFromInput(reader)
		if err := g.SelectPiece(*nextPiece); err != nil {
			fmt.Println(err.Error())
			continue
		}
		g.IncrementPlayer()
		g.Board.PrintBoard()
	}
}
