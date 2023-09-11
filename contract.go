package mango

import "fmt"

// ContractMetric represents a single position in a market.
//
// See [the Manifold API docs for GET /v0/market/marketId/positions] for more details
//
// [the Manifold API docs for GET /v0/market/marketId/positions]: https://docs.manifold.markets/api#get-v0marketmarketidpositions
type ContractMetric struct {
	ContractId    string             `json:"contractId"`
	From          map[string]Period  `json:"from,omitempty"`
	HasNoShares   bool               `json:"hasNoShares"`
	HasShares     bool               `json:"hasShares"`
	HasYesShares  bool               `json:"hasYesShares"`
	Invested      float64            `json:"invested"`
	Loan          float64            `json:"loan"`
	MaxShares     string             `json:"maxSharesOutcome,omitempty"`
	Payout        float64            `json:"payout"`
	Profit        float64            `json:"profit"`
	ProfitPercent float64            `json:"profitPercent"`
	TotalShares   map[string]float64 `json:"totalShares"`
	UserId        string             `json:"userId,omitempty"`
	UserName      string             `json:"userName,omitempty"`
	UserUsername  string             `json:"userUsername,omitempty"`
	UserAvatarUrl string             `json:"userAvatarUrl,omitempty"`
	LastBetTime   int64              `json:"lastBetTime,omitempty"`
}

// Period represents activity on a contract during a given day, week or month
type Period struct {
	Profit        float64 `json:"profit"`
	ProfitPercent float64 `json:"profitPercent"`
	Invested      float64 `json:"invested"`
	PrevValue     float64 `json:"prevValue"`
	Value         float64 `json:"value"`
}

// GetMarketPositionsRequest represents the optional parameters that can be supplied to
// get market positions via the API
//
// MarketId is required, all other fields are optional.
type GetMarketPositionsRequest struct {
	MarketID string `json:"marketId"`
	Order    string `json:"order,omitempty"`
	Top      int    `json:"top,omitempty"`
	Bottom   int    `json:"bottom,omitempty"`
	UserID   string `json:"userId,omitempty"`
}

func equalContractMetrics(cm1, cm2 ContractMetric) (bool, string) {
	// Compare each field of the objects
	if cm1.ContractId != cm2.ContractId {
		return false, fmt.Sprintf("ContractId fields are not equal: got %v expected %v", cm1.ContractId, cm2.ContractId)
	}
	if cm1.HasNoShares != cm2.HasNoShares {
		return false, fmt.Sprintf("HasNoShares fields are not equal: got %v expected %v", cm1.HasNoShares, cm2.HasNoShares)
	}
	if cm1.HasShares != cm2.HasShares {
		return false, fmt.Sprintf("HasShares fields are not equal: got %v expected %v", cm1.HasShares, cm2.HasShares)
	}
	if cm1.HasYesShares != cm2.HasYesShares {
		return false, fmt.Sprintf("HasYesShares fields are not equal: got %v expected %v", cm1.HasYesShares, cm2.HasYesShares)
	}
	if cm1.Invested != cm2.Invested {
		return false, fmt.Sprintf("Invested fields are not equal: got %v expected %v", cm1.Invested, cm2.Invested)
	}
	if cm1.Loan != cm2.Loan {
		return false, fmt.Sprintf("Loan fields are not equal: got %v expected %v", cm1.Loan, cm2.Loan)
	}
	if cm1.MaxShares != cm2.MaxShares {
		return false, fmt.Sprintf("MaxShares fields are not equal: got %v expected %v", cm1.MaxShares, cm2.MaxShares)
	}
	if cm1.Payout != cm2.Payout {
		return false, fmt.Sprintf("Payout fields are not equal: got %v expected %v", cm1.Payout, cm2.Payout)
	}
	if cm1.Profit != cm2.Profit {
		return false, fmt.Sprintf("Profit fields are not equal: got %v expected %v", cm1.Profit, cm2.Profit)
	}
	if cm1.ProfitPercent != cm2.ProfitPercent {
		return false, fmt.Sprintf("ProfitPercent fields are not equal: got %v expected %v", cm1.ProfitPercent, cm2.ProfitPercent)
	}
	if cm1.UserId != cm2.UserId {
		return false, fmt.Sprintf("UserId fields are not equal: got %v expected %v", cm1.UserId, cm2.UserId)
	}
	if cm1.UserName != cm2.UserName {
		return false, fmt.Sprintf("UserName fields are not equal: got %v expected %v", cm1.UserName, cm2.UserName)
	}
	if cm1.UserUsername != cm2.UserUsername {
		return false, fmt.Sprintf("UserUsername fields are not equal: got %v expected %v", cm1.UserUsername, cm2.UserUsername)
	}
	if cm1.UserAvatarUrl != cm2.UserAvatarUrl {
		return false, fmt.Sprintf("UserAvatarUrl fields are not equal: got %v expected %v", cm1.UserAvatarUrl, cm2.UserAvatarUrl)
	}
	if cm1.LastBetTime != cm2.LastBetTime {
		return false, fmt.Sprintf("contractId fields are not equal: got %v expected %v", cm1.LastBetTime, cm2.LastBetTime)
	}

	// Compare the 'From' maps
	if len(cm1.From) != len(cm2.From) {
		return false, fmt.Sprintf("From fields are not equal lengths: got %v expected %v", len(cm1.From), len(cm2.From))
	}
	for key, value1 := range cm1.From {
		value2, ok := cm2.From[key]
		if !ok {
			return false, fmt.Sprintf("From fields does not exist on second object: %v", cm1.From[key])
		}
		if value1 != value2 {
			return false, fmt.Sprintf("From fields are not equal: got %v expected %v", value1, value2)
		}
	}

	// Compare the 'TotalShares' maps
	if len(cm1.TotalShares) != len(cm2.TotalShares) {
		return false, fmt.Sprintf("TotalShares fields are not equal lengths: got %v expected %v", len(cm1.TotalShares), len(cm2.TotalShares))
	}
	for key, value1 := range cm1.TotalShares {
		value2, ok := cm2.TotalShares[key]
		if !ok {
			return false, fmt.Sprintf("TotalShares fields does not exist on second object: %v", cm1.TotalShares[key])
		}
		if value1 != value2 {
			return false, fmt.Sprintf("TotalShares fields are not equal: got %v expected %v", value1, value2)
		}
	}

	// If all fields are equal, the objects are equal
	return true, ""
}
