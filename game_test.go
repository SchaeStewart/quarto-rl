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