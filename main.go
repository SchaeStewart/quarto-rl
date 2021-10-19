package main

import "os"

func main() {
	game := NewGame()

	game.Loop(os.Stdin)
}
