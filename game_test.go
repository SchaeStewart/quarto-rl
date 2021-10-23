package main

import (
	"bytes"
	"testing"
)

func TestGameLoop(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("TBFR\n0\n0\nTBFC\n0\n1\nTBDR\n0\n2\nTBDC\n0\n3\n"))

	g := NewGame()
	g.Loop(&stdin)
	if g.CurrentPlayer != Player1 {
		t.Error("expected Player1 to be the winner")
	}
}

func ExampleGame() {
	var stdin bytes.Buffer
	stdin.Write([]byte("TBFR\n0\n0\nTBFC\n0\n1\nTBDR\n0\n2\nTBDC\n0\n3\n"))

	g := NewGame()
	g.Loop(&stdin)

	// Due to limitation of the golang Example tests input lines with reader input
	// seem to swallow the next line in the output

	// Output:
	// Enter your piece (T|S, B|W, F|D, R|C): Current player: 1
	// Current piece: TBFR
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Enter your X value: Enter your Y value: |TBFR|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Enter your piece (T|S, B|W, F|D, R|C): |TBFR|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Current player: 0
	// Current piece: TBFC
	// |TBFR|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Enter your X value: Enter your Y value: |TBFR|TBFC|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Enter your piece (T|S, B|W, F|D, R|C): |TBFR|TBFC|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Current player: 1
	// Current piece: TBDR
	// |TBFR|TBFC|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Enter your X value: Enter your Y value: |TBFR|TBFC|TBDR|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Enter your piece (T|S, B|W, F|D, R|C): |TBFR|TBFC|TBDR|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Current player: 0
	// Current piece: TBDC
	// |TBFR|TBFC|TBDR|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Enter your X value: Enter your Y value: |TBFR|TBFC|TBDR|TBDC|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// |0000|0000|0000|0000|
	// Winning player: 0
	//
}

// https://gocodecloud.com/blog/2018/04/01/testing-stdout-in-golang/
// TODO: add more tests
// TODO: detect tied game state
