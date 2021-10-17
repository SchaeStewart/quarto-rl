package main

import "testing"

func TestCreatePieceSet (t *testing.T) {
	set := CreatePieceSet()
	if (len(set)) != 16 {
		t.Errorf("Set should contain 16 pieces. Instead got %d", len(set))
	}
	for i, p := range set {
		for _, q := range set[i+1:] {
			if p.IsIdentical(q) {
				t.Errorf("error. duplicate pieces found p: %v, q: %v", p, q)
			}
		}
	}

}