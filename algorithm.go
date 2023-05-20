package mango

import (
	"errors"
	"fmt"
	"math"
)

// State of a given market
type State struct {
	Yes float64
	No  float64
	P   float64 // not sure what this is
}

// Adjust a probability number for a new number of YES and NO bets.
func getCPMMProbability(yes, no float64, outcome string, prob float64) float64 {
	return prob * no / ((1-prob)*yes + prob*no)
}

func SharesFromBet(state State, bet float64, outcome string) float64 {
	p := state.P
	k := math.Pow(state.Yes, p) * math.Pow(state.No, 1-p)
	switch outcome {
	case "YES":
		return state.Yes + bet - math.Pow((k*math.Pow(bet+state.No, p-1)), 1/p)
	case "NO":
		return state.No + bet - math.Pow((k*math.Pow(bet+state.Yes, -1*p)), (1/(1-p)))
	default:
		panic(fmt.Sprintf(`invalid outcome %q (should be "YES" or "NO")`, outcome))
	}
}

// Given a number of shares and an outcome, work backwards to determine what
// size bet must be made.
func BetFromShares(state State, shares, guess float64, outcome string, delta float64) (float64, error) {
	maxIterations := 100
	for i := 0; i < maxIterations; i++ {
		gotShares := SharesFromBet(state, guess, outcome)
		if math.Abs(gotShares-shares) < delta {
			return guess, nil
		}
		// I'm sure there are better search algorithms but this appears to work
		// well enough.
		guess = guess * (shares / gotShares)
	}
	return -1, errors.New("max iterations exceeded")
}

func NewBinaryProbability(state State, bet float64, outcome string) (shares float64, newProbability float64) {
	// copied from common/src/calculate-cpmm.ts:calculateCpmmPurchase
	p := state.P
	shares = SharesFromBet(state, bet, outcome)
	var newYes, newNo float64
	if outcome == "YES" {
		newYes = state.Yes - shares + bet
		newNo = state.No + bet
	} else {
		newYes = state.Yes + bet
		newNo = state.No - shares + bet
	}
	// addCpmmLiquidity
	return shares, getCPMMProbability(newYes, newNo, outcome, p)
}

// BinaryAmountToProbability returns the amount that must be bet to move the
// current market state to prob.
func BinaryAmountToProbability(state State, prob float64, outcome string) (float64, error) {
	if prob <= 0 || prob >= 1 {
		return 0, fmt.Errorf("bad probability: %v", prob)
	}
	switch outcome {
	case "NO":
		prob = 1 - prob
	case "YES":
	default:
		return 0, fmt.Errorf(`unknown outcome %q (should be "YES" or "NO")`, outcome)
	}
	p := state.P
	k := math.Pow(state.Yes, p) * math.Pow(state.No, 1-p)

	// https://www.wolframalpha.com/input?i=-1+%2B+t+-+((-1+%2B+p)+t+(k%2F(n+%2B+b))^(1%2Fp))%2Fp+solve+b
	if outcome == "YES" {
		return math.Pow((p*(prob-1))/(p-1*prob), -p) * (k - state.No*math.Pow(((p*(prob-1))/((p-1)*prob)), p)), nil
	} else {
	}
	return 0, nil
}
