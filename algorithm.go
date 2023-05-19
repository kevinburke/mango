package mango

import (
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

func NewBinaryProbability(state State, bet float64, outcome string) (shares float64, newProbability float64) {
	// copied from common/src/calculate-cpmm.ts:calculateCpmmPurchase
	p := state.P
	k := math.Pow(state.Yes, p) * math.Pow(state.No, 1-p)
	switch outcome {
	case "YES":
		shares = state.Yes + bet - math.Pow((k*math.Pow(bet+state.No, p-1)), 1/p)
	case "NO":
		shares = state.No + bet - math.Pow((k*math.Pow(bet+state.Yes, -1*p)), (1/(1-p)))
	default:
		panic(fmt.Sprintf(`invalid outcome %q (should be "YES" or "NO")`, outcome))
	}
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
	}
	return 0, nil
}
