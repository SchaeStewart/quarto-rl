package main

import (
	"fmt"
	"time"
)

func main() {
	game := NewGame()
	// game.Loop(os.Stdin)
	agent1 := NewRandomAgent(time.Now().UnixNano(), game)
	agent2 := NewRandomAgent(time.Now().UnixNano() / 2, game)
	PlayAgents(agent1, agent2, game)
}

func PlayAgents (agent1, agent2 Agent, game *Game) {
	agent1.SelectPiece()
	for !game.IsFinished() {
		x, y := agent2.PlayPiece()
		game.PlayPiece(x, y)
		game.SelectPiece(agent2.SelectPiece())
		game.Print()
		if game.IsFinished() {
			continue
		}
		x, y = agent1.PlayPiece()
		game.PlayPiece(x, y)
		game.SelectPiece(agent1.SelectPiece())
		game.Print()
	}
	fmt.Println("Winning player:", game.CurrentPlayer)

}
