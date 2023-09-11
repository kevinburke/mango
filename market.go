package mango

import (
	"fmt"
	"reflect"
)

// OutcomeType represents the different types of markets
// available on Manifold.
type OutcomeType string

const (
	Binary         OutcomeType = "BINARY"
	FreeResponse   OutcomeType = "FREE_RESPONSE"
	MultipleChoice OutcomeType = "MULTIPLE_CHOICE"
	Numeric        OutcomeType = "NUMERIC"
	PseudoNumeric  OutcomeType = "PSEUDO_NUMERIC"
)

// Pool represents the potential outcomes for a market.
type Pool struct {
	No       float64 `json:"NO,omitempty"`
	Yes      float64 `json:"YES,omitempty"`
	Option0  float64 `json:"0,omitempty"`
	Option1  float64 `json:"1,omitempty"`
	Option2  float64 `json:"2,omitempty"`
	Option3  float64 `json:"3,omitempty"`
	Option4  float64 `json:"4,omitempty"`
	Option5  float64 `json:"5,omitempty"`
	Option6  float64 `json:"6,omitempty"`
	Option7  float64 `json:"7,omitempty"`
	Option8  float64 `json:"8,omitempty"`
	Option9  float64 `json:"9,omitempty"`
	Option10 float64 `json:"10,omitempty"`
	Option11 float64 `json:"11,omitempty"`
	Option12 float64 `json:"12,omitempty"`
	Option13 float64 `json:"13,omitempty"`
	Option14 float64 `json:"14,omitempty"`
	Option15 float64 `json:"15,omitempty"`
	Option16 float64 `json:"16,omitempty"`
	Option17 float64 `json:"17,omitempty"`
	Option18 float64 `json:"18,omitempty"`
	Option19 float64 `json:"19,omitempty"`
}

// Answer represents a potential answer on a free response market
type Answer struct {
	Id          string  `json:"id"`
	Username    string  `json:"username"`
	Name        string  `json:"name"`
	UserId      string  `json:"userId"`
	CreatedTime int64   `json:"createdTime"`
	AvatarUrl   string  `json:"avatarUrl"`
	Number      int64   `json:"number"`
	ContractId  string  `json:"contractId"`
	Text        string  `json:"text"`
	Probability float64 `json:"probability"`
}

// GetMarketsRequest represents the optional parameters that can be supplied to
// get markets via the API
type GetMarketsRequest struct {
	Before string `json:"before,omitempty"`
	Limit  int64  `json:"limit,omitempty"`
}

// PostMarketRequest represents the parameters required to create a new market via the API
type PostMarketRequest struct {
	OutcomeType         OutcomeType `json:"outcomeType"`
	Question            string      `json:"question"`
	Description         string      `json:"description,omitempty"`
	DescriptionHtml     string      `json:"descriptionHtml,omitempty"`
	DescriptionMarkdown string      `json:"descriptionMarkdown,omitempty"`
	CloseTime           int64       `json:"closeTime,omitempty"`
	Visibility          string      `json:"visibility,omitempty"`
	GroupIds            []string    `json:"groupIds,omitempty"`
	InitialProb         int64       `json:"initialProb,omitempty"`
	Min                 int64       `json:"min,omitempty"`
	Max                 int64       `json:"max,omitempty"`
	IsLogScale          bool        `json:"isLogScale,omitempty"`
	InitialVal          int64       `json:"initialValue,omitempty"`
	Answers             []string    `json:"answers,omitempty"`
}

// ResolveMarketRequest represents the parameters required to resolve a market via the API
type ResolveMarketRequest struct {
	Outcome        string       `json:"outcome"`
	Resolutions    []Resolution `json:"resolutions,omitempty"`
	ProbabilityInt int64        `json:"probabilityInt,omitempty"`
	Value          float64      `json:"value,omitempty"`
}

// Resolution represents the percentage a given answer should resolve to
// on a market
type Resolution struct {
	Answer int64 `json:"answer"`
	Pct    int64 `json:"pct"`
}

// SellSharesRequest represents a request to sell shares
type SellSharesRequest struct {
	Outcome string `json:"outcome,omitempty"`
	Shares  int64  `json:"shares,omitempty"`
}

type marketIdResponse struct {
	Id string `json:"id"`
}

type Market interface {
	LiteMarket | FullMarket
}

// LiteMarket represents a LiteMarket object in the Manifold backend.
// A LiteMarket is similar to a [FullMarket], except it has fewer fields.
//
// See [the Manifold API docs for GET /v0/markets] for more details
//
// [the Manifold API docs for GET /v0/markets]: https://docs.manifold.markets/api#get-v0markets
type LiteMarket struct {
	Id                    string        `json:"id"`
	CreatorId             string        `json:"creatorId"`
	CreatorUsername       string        `json:"creatorUsername"`
	CreatorName           string        `json:"creatorName"`
	CreatedTime           int64         `json:"createdTime"`
	CreatorAvatarUrl      string        `json:"creatorAvatarUrl"`
	CloseTime             int64         `json:"closeTime"`
	Question              string        `json:"question"`
	Tags                  []interface{} `json:"tags"`
	Url                   string        `json:"url"`
	Pool                  Pool          `json:"pool,omitempty"`
	Probability           float64       `json:"probability,omitempty"`
	P                     float64       `json:"p,omitempty"`
	TotalLiquidity        float64       `json:"totalLiquidity,omitempty"`
	OutcomeType           OutcomeType   `json:"OutcomeType"`
	Mechanism             string        `json:"mechanism"`
	Volume                float64       `json:"volume"`
	Volume24Hours         float64       `json:"volume24Hours"`
	IsResolved            bool          `json:"isResolved"`
	LastUpdatedTime       int64         `json:"lastUpdatedTime,omitempty"`
	Min                   float64       `json:"min,omitempty"`
	Max                   float64       `json:"max,omitempty"`
	IsLogScale            bool          `json:"isLogScale,omitempty"`
	Resolution            string        `json:"resolution,omitempty"`
	ResolutionTime        int64         `json:"resolutionTime,omitempty"`
	ResolutionProbability float64       `json:"resolutionProbability,omitempty"`
}

// FullMarket represents a FullMarket object in the Manifold backend.
// A FullMarket is similar to a [LiteMarket], except it has more fields.
//
// See [the Manifold API docs for GET /v0/market/[marketId]] for more details
//
// [the Manifold API docs for GET /v0/market/[marketId]]: https://docs.manifold.markets/api#get-v0marketmarketid
type FullMarket struct {
	Id                    string      `json:"id"`
	CreatorId             string      `json:"creatorId"`
	CreatorUsername       string      `json:"creatorUsername"`
	CreatorName           string      `json:"creatorName"`
	CreatedTime           int64       `json:"createdTime"`
	CreatorAvatarUrl      string      `json:"creatorAvatarUrl"`
	CloseTime             int64       `json:"closeTime"`
	Question              string      `json:"question"`
	Answers               []Answer    `json:"answers,omitempty"`
	Tags                  []string    `json:"tags"`
	Url                   string      `json:"url"`
	Pool                  Pool        `json:"pool"`
	Probability           float64     `json:"probability"`
	P                     float64     `json:"p"`
	TotalLiquidity        float64     `json:"totalLiquidity"`
	OutcomeType           OutcomeType `json:"OutcomeType"`
	Mechanism             string      `json:"mechanism"`
	Volume                float64     `json:"volume"`
	Volume24Hours         float64     `json:"volume24Hours"`
	IsResolved            bool        `json:"isResolved"`
	Resolution            string      `json:"resolution"`
	ResolutionTime        int64       `json:"resolutionTime"`
	ResolutionProbability float64     `json:"resolutionProbability"`
	LastUpdatedTime       int64       `json:"lastUpdatedTime"`
	// Description field returns HTML marshalled to JSON, see https://tiptap.dev/guide/output#option-1-json
	// Description     string `json:"description"` TODO: work out how to parse this field
	TextDescription string `json:"textDescription"`
}

func equalFullMarkets(m1, m2 FullMarket) (bool, string) {
	if m1.Id != m2.Id {
		return false, fmt.Sprintf("Id fields are not equal: got %v expected %v", m1.Id, m2.Id)
	}
	if m1.CreatorId != m2.CreatorId {
		return false, fmt.Sprintf("CreatorId fields are not equal: got %v expected %v", m1.CreatorId, m2.CreatorId)
	}
	if m1.CreatorUsername != m2.CreatorUsername {
		return false, fmt.Sprintf("CreatorUsername fields are not equal: got %v expected %v", m1.CreatorUsername, m2.CreatorUsername)
	}
	if m1.CreatorName != m2.CreatorName {
		return false, fmt.Sprintf("CreatorName fields are not equal: got %v expected %v", m1.CreatorName, m2.CreatorName)
	}
	if m1.CreatedTime != m2.CreatedTime {
		return false, fmt.Sprintf("CreatedTime fields are not equal: got %v expected %v", m1.CreatedTime, m2.CreatedTime)
	}
	if m1.CreatorAvatarUrl != m2.CreatorAvatarUrl {
		return false, fmt.Sprintf("CreatorAvatarUrl fields are not equal: got %v expected %v", m1.CreatorAvatarUrl, m2.CreatorAvatarUrl)
	}
	if m1.CloseTime != m2.CloseTime {
		return false, fmt.Sprintf("CloseTime fields are not equal: got %v expected %v", m1.CloseTime, m2.CloseTime)
	}
	if m1.Question != m2.Question {
		return false, fmt.Sprintf("Question fields are not equal: got %v expected %v", m1.Question, m2.Question)
	}
	if !reflect.DeepEqual(m1.Tags, m2.Tags) {
		return false, fmt.Sprintf("Tags fields are not equal: got %v expected %v", m1.Tags, m2.Tags)
	}
	if m1.Url != m2.Url {
		return false, fmt.Sprintf("Url fields are not equal: got %v expected %v", m1.Url, m2.Url)
	}
	if !reflect.DeepEqual(m1.Pool, m2.Pool) {
		return false, fmt.Sprintf("Pool fields are not equal: got %v expected %v", m1.Pool, m2.Pool)
	}

	return true, ""
}
