package mango

import (
	"fmt"
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

func TestWithFee(t *testing.T) {
	// Trump win 7: mkt1 0.57582, mkt2 0.56258 (diff of 0.01324), arb size: 35.1054
	// 2024/07/23 05:42:27 INFO completed bet market=2 outcome=YES shares=251.24850038880322 amount_bet=146.0946 bet_size=146.094552 before_prob=0.562576 after_prob=0.565948
	// 2024/07/23 05:42:27 INFO completed bet market=1 outcome=NO shares=249.1961050957325 amount_bet=111 bet_size=111 before_prob=0.575816 after_prob=0.567596
	// 2024/07/23 05:42:27 INFO arbitrage complete name="Trump win 7" spent=257.09455219478264 min_profit_percent=-3.0722% min_profit=-7.898447099050145 mkt1_shares=249.19611 mkt2_shares=251.2485 latency=293ms
	// shares1 249.1961050957325 or shares2 251.24850038880322 less than spent 257.09455219478264, no arbitrage (orig state1: mango.State{Yes:5024.771077726842, No:11420.508230720294, P:0.3739258597372285}, state2: mango.State{Yes:23333.02380679098, No:15663.17014467531, P:0.6570513944641774})

	state1 := State{Yes: 5024.771077726842, No: 11420.508230720294, P: 0.3739258597372285}
	// state2 := State{Yes: 23333.02380679098, No: 15663.17014467531, P: 0.6570513944641774}

	shares := state1.Bet(146.0946, "YES")
	fmt.Println("got shares", shares)
	t.Fail()
}
