package main

import "testing"

func TestCreatePieceSet(t *testing.T) {
	set := CreatePieceSet()
	if (len(set)) != 16 {
		t.Errorf("Set should contain 16 pieces. Instead got %d", len(set))
	}
	for i, p := range set {
		for _, q := range set[i+1:] {
			if IsIdentical(p, q) {
				t.Errorf("error. duplicate pieces found p: %v, q: %v", p, q)
			}
		}
	}
}

func TestPieceString(t *testing.T) {
	table := []struct {
		p        *Piece
		expected string
	}{
		{
			p:        &Piece{true, true, true, true},
			expected: "TBFR",
		},
		{
			p:        nil,
			expected: "0000",
		},
		{
			p:        &Piece{},
			expected: "SWDC",
		},
	}

	for _, tt := range table {
		result := tt.p.String()
		if result != tt.expected {
			t.Errorf("piece not converted to string. Expected %s, instead got %s", tt.expected, result)
		}
	}
}
