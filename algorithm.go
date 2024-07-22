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

// Returns the new probability
func (s *State) Bet(bet float64, outcome string) float64 {
	if s == nil {
		panic("mango.State.Bet: nil *State")
	}
	shares := SharesFromBet(*s, bet, outcome)
	switch outcome {
	case "YES":
		s.Yes = s.Yes - shares + bet
		s.No = s.No + bet
	case "NO":
		s.Yes = s.Yes + bet
		s.No = s.No - shares + bet
	default:
		panic(fmt.Sprintf("unknown outcome value %q", outcome))
	}
	fee := float64(0)
	prob := getCPMMProbability(s.Yes, s.No, s.P)
	numerator := prob * (fee + s.Yes)
	denominator := fee - s.No*(prob-1) + prob*s.Yes
	s.P = numerator / denominator
	return shares
}

func (s State) Probability() float64 {
	return getCPMMProbability(s.Yes, s.No, s.P)
}

func getCPMMProbability(yes, no float64, p float64) float64 {
	return (p * no) / ((1-p)*yes + p*no)
}

/*
from fees.ts, July 2024
const TAKER_FEE_CONSTANT = 0.07
export const getTakerFee = (shares: number, prob: number) => {
  return TAKER_FEE_CONSTANT * prob * (1 - prob) * shares
}
*/

const FeeConstant = 0.07

func getTakerFee(shares, prob float64) float64 {
	return FeeConstant * prob * (1 - prob) * shares
}

func SharesFromBet(state State, bet float64, outcome string) float64 {
	// calculate shares
	// calculate fee
	// subtract fee from bet
	// calculate remaining shares
	fee := 0.0
	// this is just copied from getCpmmFee
	for i := 0; i < 10; i++ {
		betAmountAfterFee := bet - fee
		sharesAfterFee := sharesFromBet(state, betAmountAfterFee, outcome)
		averageProb := betAmountAfterFee / sharesAfterFee
		fee = getTakerFee(averageProb, sharesAfterFee)
	}
	remainingBet := bet - fee
	return sharesFromBet(state, remainingBet, outcome)
}

func sharesFromBet(state State, bet float64, outcome string) float64 {
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

func NewBinaryProbability(state *State, bet float64, outcome string) (shares float64, newProbability float64) {
	shares = state.Bet(bet, outcome)
	return shares, state.Probability()
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
