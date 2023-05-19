package mango

import "testing"

var binaryProbTestCases = []struct {
	Yes, No    float64
	P          float64
	Bet        float64
	Outcome    string
	WantShares float64
	WantP      float64
}{
	{7692.807926448773, 185.31845682230974, 0.10882454509729465, 10, "YES", 2700.484088026393, 0.004745361855569013},
}

func TestBinaryProbability(t *testing.T) {
	for _, tt := range binaryProbTestCases {
		state := state{
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
