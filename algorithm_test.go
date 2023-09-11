package mango

import (
	"math"
	"testing"
)

var binaryProbTestCases = []struct {
	Yes, No    float64
	P          float64
	Bet        float64
	Outcome    string
	WantShares float64
	WantP      float64
}{
	{7692.807926448773, 185.31845682230974, 0.10882454509729465, 10, "YES", 2700.484088026393, 0.004745361855569013},
	{7692.807926448773, 185.31845682230974, 0.10882454509729465, 10, "YES", 2700.484088026393, 0.004745361855569013},
}

func TestBinaryProbability(t *testing.T) {
	for _, tt := range binaryProbTestCases {
		state := &State{
			Yes: tt.Yes,
			No:  tt.No,
			P:   tt.P,
		}
		gotShares, gotP := NewBinaryProbability(state, tt.Bet, tt.Outcome)
		if gotShares != tt.WantShares {
			t.Errorf("NewBinaryProbability(%#v, %v, %q): got %v shares, want %v shares", state, tt.Bet, tt.Outcome, gotShares, tt.WantShares)
		}
		if gotP != tt.WantP {
			t.Errorf("NewBinaryProbability(%#v, %v, %q): got %v prob, want %v prob", state, tt.Bet, tt.Outcome, gotP, tt.WantP)
		}
	}
}

func TestBetFromShares(t *testing.T) {
	testCaseAdditions := []struct {
		Delta float64
		Guess float64
	}{
		{0.02, 5},
		{0.02, 15},
	}
	for i, tt := range binaryProbTestCases {
		state := State{
			Yes: tt.Yes,
			No:  tt.No,
			P:   tt.P,
		}
		guess := testCaseAdditions[i].Guess
		delta := testCaseAdditions[i].Delta
		gotBet, err := BetFromShares(state, tt.WantShares, guess, tt.Outcome, delta)
		if err != nil {
			t.Fatal(err)
		}
		if math.Abs(gotBet-tt.Bet) >= delta {
			t.Errorf("BetFromShares(%#v, %v, %v, %q): got %v shares, want %v shares", state, tt.WantShares, delta, tt.Outcome, gotBet, tt.Bet)
		}
	}
}
