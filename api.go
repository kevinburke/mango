// Package mango provides a client that can be used to make calls
// to the Manifold Markets API.
//
// Currently, it provides wrapper functions for every documented
// API call that Manifold offers. It also offers data types representing
// each data structure that can be returned by the API.
//
// See [the Manifold API docs] for more details.
//
// [the Manifold API docs]: https://docs.manifold.markets/api
package mango

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

// TODO: POST doc comments need a note that they won't work without an API key

// GetAuthenticatedUser returns the [User] associated with the current API key.
// If the key is invalid or not present, returns nil and an error.
//
// See [the Manifold API docs for GET /v0/me] for more details.
//
// [the Manifold API docs for GET /v0/me]: https://docs.manifold.markets/api#get-v0me
func (mc *Client) GetAuthenticatedUser(ctx context.Context) (*User, error) {
	resp, err := mc.makeRequest(ctx, http.MethodGet, requestURL(mc.url, getMe, "", ""), nil)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, User{})
}

// getBets returns a slice of [Bet] and an error. It takes a [GetBetsRequest] which has the following
// optional parameters:
//   - [GetBetsRequest.UserId]
//   - [GetBetsRequest.Username]
//   - [GetBetsRequest.ContractId]
//   - [GetBetsRequest.ContractSlug]
//   - [GetBetsRequest.Before] - the ID of the bet before which the list will start.
//   - [GetBetsRequest.Limit] - the maximum and default limit is 1000.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/bets] for more details.
//
// [the Manifold API docs for GET /v0/bets]: https://docs.manifold.markets/api#get-v0bets
func (mc *Client) GetBets(gbr GetBetsRequest) (*[]Bet, error) {
	if gbr.Limit == 0 {
		gbr.Limit = defaultLimit
	}
	resp, err := mc.client.Get(requestURL(
		mc.url,
		getBets, "", "",
		"userId", gbr.UserId,
		"username", gbr.Username,
		"contractId", gbr.ContractId,
		"contractSlug", gbr.ContractSlug,
		"before", gbr.Before,
		"limit", strconv.FormatInt(gbr.Limit, 10),
	))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, []Bet{})
}

// getComments returns a slice of [Comment] and an error. It takes a [GetCommentsRequest] which has the following
// optional parameters:
//   - [GetCommentsRequest.ContractId]
//   - [GetCommentsRequest.ContractSlug]
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/comments] for more details.
//
// [the Manifold API docs for GET /v0/comments]: https://docs.manifold.markets/api#get-v0comments
func (mc *Client) GetComments(gcr GetCommentsRequest) (*[]Comment, error) {
	if gcr.ContractId == "" && gcr.ContractSlug == "" {
		return nil, fmt.Errorf("either contractID or contractSlug must be specified")
	}

	resp, err := mc.client.Get(requestURL(mc.url, getComments, "", "",
		"contractId", gcr.ContractId,
		"contractSlug", gcr.ContractSlug,
	))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, []Comment{})
}

// GetGroupById returns a [Group] by its unique id.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/group/by-id/id] for more details.
//
// [the Manifold API docs for GET /v0/group/by-id/id]: https://docs.manifold.markets/api#get-v0groupby-idid
func (mc *Client) GetGroupById(id string) (*Group, error) {
	resp, err := mc.client.Get(requestURL(mc.url, getGroupByID, id, ""))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, Group{})
}

// getGroupBySlug returns a [Group] by its slug.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/group/slug] for more details.
//
// [the Manifold API docs for GET /v0/group/slug]: https://docs.manifold.markets/api#get-v0groupslug
func (mc *Client) GetGroupBySlug(slug string) (*Group, error) {
	resp, err := mc.client.Get(requestURL(mc.url, getGroupBySlug, slug, ""))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, Group{})
}

// getGroups returns a slice of [Group]. Optionally a userId can be passed:
// in this case, only groups available to the given user will be returned.
// Results are unordered.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/groups] for more details.
//
// [the Manifold API docs for GET /v0/groups]: https://docs.manifold.markets/api#get-v0groups
func (mc *Client) GetGroups(userId *string) (*[]Group, error) {
	uid := ""
	if userId != nil {
		uid = *userId
	}
	resp, err := mc.client.Get(requestURL(
		mc.url, getGroups, "", "",
		"availableToUserId", uid,
	))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, []Group{})
}

// getMarketByID returns a [FullMarket] by its unique id.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/market/marketId] for more details.
//
// [the Manifold API docs for GET /v0/market/marketId]: https://docs.manifold.markets/api#get-v0marketmarketid
func (mc *Client) GetMarketByID(ctx context.Context, id string) (*FullMarket, error) {
	resp, err := mc.makeRequest(ctx, http.MethodGet, requestURL(mc.url, getMarketByID, id, ""), nil)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, FullMarket{})
}

// getMarketBySlug returns a [FullMarket] by its unique slug.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/slug/marketSlug] for more details.
//
// [the Manifold API docs for GET /v0/slug/marketSlug]: https://docs.manifold.markets/api#get-v0slugmarketslug
func (mc *Client) GetMarketBySlug(slug string) (*FullMarket, error) {
	resp, err := mc.client.Get(requestURL(mc.url, getMarketBySlug, slug, ""))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, FullMarket{})
}

// getMarkets returns a slice of [LiteMarket] and an error. It takes a [GetMarketsRequest] which has the following
// optional parameters:
//   - [GetMarketsRequest.Before] - the ID of the market before which the list will start.
//   - [GetMarketsRequest.Limit] - the maximum and default limit is 1000.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/markets] for more details.
//
// [the Manifold API docs for GET /v0/markets]: https://docs.manifold.markets/api#get-v0markets
func (mc *Client) GetMarkets(ctx context.Context, gmr GetMarketsRequest) (*[]LiteMarket, error) {
	if gmr.Limit == 0 {
		gmr.Limit = defaultLimit
	}

	url := requestURL(
		mc.url, getMarkets,
		"",
		"",
		"limit", strconv.FormatInt(gmr.Limit, 10), "before", gmr.Before,
	)
	resp, err := mc.makeRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, []LiteMarket{})
}

// GetMarketsForGroup returns a slice of [LiteMarket] and an error. It takes a group ID to retrieve the markets for.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/group/by-id/id/markets] for more details
//
// [the Manifold API docs for GET /v0/group/by-id/id/markets]: https://docs.manifold.markets/api#get-v0groupby-ididmarkets
func (mc *Client) GetMarketsForGroup(id string) (*[]LiteMarket, error) {
	resp, err := mc.client.Get(requestURL(mc.url, getGroupByID, id, marketsSuffix))
	if err != nil {
		log.Printf("error making http request: %w", err)
	}

	return parseResponse(resp, []LiteMarket{})
}

// GetMarketPositions returns positions for a given [Market] ID. It takes a [GetMarketPositionsRequest]
// which has the following parameters:
//   - [GetMarketPositionsRequest.MarketId] - Required.
//   - [GetMarketPositionsRequest.Order] - Optional. The field to order results by. Default: "profit". Options: "shares" or "profit". // TODO: make this an enum
//   - [GetMarketPositionsRequest.Top] - Optional. The number of top positions (ordered by order) to return. Default: null.
//   - [GetMarketPositionsRequest.Bottom] - Optional. The number of bottom positions (ordered by order) to return. Default: null.
//   - [GetMarketPositionsRequest.UserId] - Optional. The user ID to query by. Default: null. If provided, only the position for this user will be returned.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/market/marketId/positions] for more details.
//
// [the Manifold API docs for GET /v0/market/marketId/positions]: https://docs.manifold.markets/api#get-v0marketmarketidpositions
func (mc *Client) GetMarketPositions(ctx context.Context, gmpr GetMarketPositionsRequest) (*[]ContractMetric, error) {
	if gmpr.MarketID == "" {
		return nil, fmt.Errorf("no market ID provided")
	}

	var t, b string
	if gmpr.Top != 0 {
		t = strconv.Itoa(gmpr.Top)
	} else {
		t = "null"
	}

	if gmpr.Bottom != 0 {
		b = strconv.Itoa(gmpr.Bottom)
	} else {
		b = "null"
	}

	url := requestURL(mc.url, getMarketByID, gmpr.MarketID, positionsSuffix,
		"order", gmpr.Order,
		"top", t,
		"bottom", b,
		"userId", gmpr.UserID,
	)
	fmt.Println("request URL", url)
	resp, err := mc.makeRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, []ContractMetric{})
}

// SearchMarkets returns a slice of [FullMarket] that are the results of searching for provided terms.
// It can take any number of strings. Each string should be an individual search term.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/search-markets] for more details.
//
// [the Manifold API docs for GET /v0/search-markets]: https://docs.manifold.markets/api#get-v0search-markets
func (mc *Client) SearchMarkets(terms ...string) (*[]FullMarket, error) {
	ts := ""

	for i, t := range terms {
		ts += strings.TrimSpace(t)
		if i+1 < len(terms) {
			ts += " "
		}
	}

	resp, err := mc.client.Get(requestURL(mc.url, getSearchMarkets, "", "",
		"terms", ts,
	))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, []FullMarket{})
}

// getUserByID returns a [User] by user id.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/user/by-id/id] for more details.
//
// [the Manifold API docs for GET /v0/user/by-id/id]: https://docs.manifold.markets/api#get-v0userby-idid
func (mc *Client) GetUserByID(id string) (*User, error) {
	resp, err := mc.client.Get(requestURL(mc.url, getUserByID, id, ""))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, User{})
}

// getUserByUsername returns a [User] by username
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/user/username] for more details.
//
// [the Manifold API docs for GET /v0/user/username]: https://docs.manifold.markets/api#get-v0userusername
func (mc *Client) GetUserByUsername(un string) (*User, error) {
	resp, err := mc.client.Get(requestURL(mc.url, getUserByUsername, un, ""))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, User{})
}

// getUsers returns a slice of [User] and an error. It takes a [GetUsersRequest] which has the following parameters:
//   - [GetUsersRequest.Before] - Optional. The ID of the user before which the list will start.
//   - [GetUsersRequest.Limit] - Optional. The default and maximum limit is 1000.
//
// If there is an error making the request, then nil and an error
// will be returned.
//
// See [the Manifold API docs for GET /v0/markets] for more details.
//
// [the Manifold API docs for GET /v0/markets]: https://docs.manifold.markets/api#get-v0markets
func (mc *Client) GetUsers(gur GetUsersRequest) (*[]User, error) {
	if gur.Limit == 0 {
		gur.Limit = defaultLimit
	}

	resp, err := mc.client.Get(requestURL(
		mc.url, getUsers,
		"",
		"",
		"limit", strconv.FormatInt(gur.Limit, 10), "before", gur.Before,
	))
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	return parseResponse(resp, []User{})
}

// postBet makes a new bet on a market. It takes a [PostBetRequest] which has the following parameters:
//   - [PostBetRequest.Amount] - Required.
//   - [PostBetRequest.ContractId] - Required.
//   - [PostBetRequest.Outcome] - Required.
//   - [PostBetRequest.LimitProb] - Optional. A number between 0.001 and 0.999 inclusive representing the limit probability for your bet
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/bet] for more details.
//
// [the Manifold API docs for POST /v0/bet]: https://docs.manifold.markets/api#post-v0bet
func (mc *Client) PostBet(ctx context.Context, pbr PostBetRequest) (*BetResponse, error) {
	jsonBody, err := json.Marshal(pbr)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	resp, err := mc.makeRequest(ctx, http.MethodPost, requestURL(
		mc.url, postBet,
		"",
		""), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("mango: could not create bet: %w", err)
	}
	return parseResponse(resp, BetResponse{})
}

type BetResponse struct {
	ContractID        string  `json:"contractId"`
	BetID             string  `json:"betId"`
	IsFilled          bool    `json:"isFilled"`
	IsCancelled       bool    `json:"isCancelled"`
	Amount            float64 `json:"amount"`
	OrderAmount       float64 `json:"orderAmount"`
	Shares            float64 `json:"shares"`
	ProbabilityBefore float64 `json:"probBefore"`
	ProbabilityAfter  float64 `json:"probAfter"`
	Visibility        string  `json:"visibility"`
	Outcome           string  `json:"outcome"`
	LoanAmount        float64 `json:"loanAmount"`
	CreatedTime       int64   `json:"createdTime"`
}

// CancelBet cancels an existing limit order for the given betId.
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/bet/cancel/id] for more details.
//
// [the Manifold API docs for POST /v0/bet/cancel/id]: https://docs.manifold.markets/api#post-v0betcancelid
func (mc *Client) CancelBet(ctx context.Context, betId string) error {
	resp, err := mc.makeRequest(ctx, http.MethodPost, requestURL(
		mc.url, postCancellation,
		betId,
		""), nil)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	_, err = parseResponse(resp, struct{}{})
	return err
}

// CreateMarket creates a new market. It takes a [PostMarketRequest] which has the following parameters:
//   - [PostMarketRequest.OutcomeType] - Required. One of [OutcomeType]
//   - [PostMarketRequest.Question] - Required.
//   - [PostMarketRequest.Description] - Optional. Non-rich text description.
//   - [PostMarketRequest.DescriptionHtml] - Optional.
//   - [PostMarketRequest.DescriptionMarkdown] - Optional.
//   - [PostMarketRequest.CloseTime] - Optional. Milliseconds since the epoch. Default 7 days from time of creation.
//   - [PostMarketRequest.Visibility] - Optional. One of "public" or "unlisted" TODO: make this an enum
//   - [PostMarketRequest.GroupId] - Optional. A group to show the market under.
//   - [PostMarketRequest.InitialProb] - Required for binary markets. Must be between 1 and 99.
//   - [PostMarketRequest.Min] - Required for numeric markets. The minimum value that the market may resolve to.
//   - [PostMarketRequest.Max] - Required for numeric markets. The maximum value that the market may resolve to.
//   - [PostMarketRequest.IsLogScale] - Required for numeric markets. If true, your numeric market will increase exponentially from min to max.
//   - [PostMarketRequest.InitialVal] - Required for numeric markets. An initial value for the market, between min and max, exclusive.
//   - [PostMarketRequest.Answers] - Required for multiple choice markets. An array of strings, each of which will be a valid answer for the market.
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/market] for more details.
//
// [the Manifold API docs for POST /v0/market]: https://docs.manifold.markets/api#post-v0market
func (mc *Client) CreateMarket(ctx context.Context, pmr PostMarketRequest) (*string, error) {
	// TODO: add input validation
	jsonBody, err := json.Marshal(pmr)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	resp, err := mc.makeRequest(ctx, http.MethodPost, requestURL(
		mc.url, postMarket,
		"",
		""), bodyReader)

	if resp.StatusCode != http.StatusOK {
		// TODO: print the response body here too
		return nil, fmt.Errorf("market creation failed with status %d", resp.StatusCode)
	}

	mir, err := parseResponse(resp, marketIdResponse{})
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return &mir.Id, nil
}

// AddLiquidity adds a given amount of liquidity to a given market.
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/market/marketId/add-liquidity] for more details.
//
// [the Manifold API docs for POST /v0/market/marketId/add-liquidity]: https://docs.manifold.markets/api#post-v0marketmarketidadd-liquidity
func (mc *Client) AddLiquidity(marketId string, amount int64) error {
	amt := struct {
		Amount int64 `json:"amount"`
	}{amount}

	jsonBody, err := json.Marshal(amt)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	resp, err := mc.makeRequest(context.TODO(), http.MethodPost, requestURL(
		mc.url, postMarket,
		marketId,
		liquiditySuffix), bodyReader)
	if err != nil {
		return fmt.Errorf("client: error making http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("liquidity addition failed with status %d", resp.StatusCode)
	}

	return nil
}

// CloseMarket updates the closing time of a given market to be the given epoch timestamp,
// or closes it immediately if no new time is provided.
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/market/marketId/close] for more details.
//
// [the Manifold API docs for POST /v0/market/marketId/close]: https://docs.manifold.markets/api#post-v0marketmarketidclose
func (mc *Client) CloseMarket(marketId string, ct *int64) error {
	if ct == nil {
		ct = new(int64)
	}

	c := struct {
		CloseTime int64 `json:"closeTime,omitempty"`
	}{*ct}

	jsonBody, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL(
		mc.url, postMarket,
		marketId,
		closureSuffix), bodyReader)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}

	resp, err := mc.postRequest(req)
	if err != nil {
		return fmt.Errorf("client: error making http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("market closure failed with status %d", resp.StatusCode)
	}

	return nil
}

// AddMarketToGroup adds a given market to a given group.
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/market/marketId/group] for more details.
//
// [the Manifold API docs for POST /v0/market/marketId/group]: https://docs.manifold.markets/api#post-v0marketmarketidgroup
func (mc *Client) AddMarketToGroup(ctx context.Context, marketId, groupId string) error {
	g := struct {
		GroupId string `json:"groupId,omitempty"`
	}{groupId}

	jsonBody, err := json.Marshal(g)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	resp, err := mc.makeRequest(ctx, http.MethodPost, requestURL(
		mc.url, postMarket,
		marketId,
		"/group"), bodyReader)

	if err != nil {
		return fmt.Errorf("client: error making http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("adding market to group failed with status %d", resp.StatusCode)
	}

	return nil
}

// TODO: add more information about input params to this comment

// ResolveMarket creates a new market. It takes a [ResolveMarketRequest] which has the following parameters:
//   - [ResolveMarketRequest.Outcome] - Required. One of "YES", "NO", "MKT", "CANCEL" or a number depending on market type. TODO: make this an enum
//   - [ResolveMarketRequest.Resolutions] - Required for free response or multiple choice markets.
//   - [ResolveMarketRequest.ProbabilityInt] - Required if value is present.
//   - [ResolveMarketRequest.Value] - Optional, only relevant to numeric markets.
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/market/marketId/resolve] for more details.
//
// [the Manifold API docs for POST /v0/market/marketId/resolve]: https://docs.manifold.markets/api#post-v0marketmarketidresolve
func (mc *Client) ResolveMarket(marketId string, rmr ResolveMarketRequest) error {
	jsonBody, err := json.Marshal(rmr)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL(
		mc.url, postMarket,
		marketId,
		resolutionSuffix), bodyReader)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}

	resp, err := mc.postRequest(req)
	if err != nil {
		return fmt.Errorf("client: error making http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("market resolution failed with status %d", resp.StatusCode)
	}

	return nil
}

// SellShares creates a new market. It takes a [SellSharesRequest] which has the following parameters:
//   - [SellSharesRequest.Outcome] - Optional. One of "YES" or "NO". If omitted and only one kind of shares are held, sells those. TODO: make this an enum
//   - [SellSharesRequest.Shares] - Optional. If omitted, all shares held will be sold.
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/market/marketId/sell] for more details.
//
// [the Manifold API docs for POST /v0/market/marketId/sell]: https://docs.manifold.markets/api#post-v0marketmarketidsell
func (mc *Client) SellShares(marketId string, ssr SellSharesRequest) error {
	jsonBody, err := json.Marshal(ssr)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL(
		mc.url, postMarket,
		marketId,
		sellSuffix), bodyReader)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}

	resp, err := mc.postRequest(req)
	if err != nil {
		return fmt.Errorf("client: error making http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("selling shares failed with status %d", resp.StatusCode)
	}

	return nil
}

// postComment makes a new bet on a market. It takes a [PostCommentRequest] which has the following parameters:
//   - [PostCommentRequest.ContractId] - Required.
//   - [PostCommentRequest.Content] - Optional. A plaintext string.
//   - [PostCommentRequest.Html] - Optional.
//   - [PostCommentRequest.Markdown] - Optional.
//
// If there is an error making the request, then an error will be returned.
//
// See [the Manifold API docs for POST /v0/comment] for more details.
//
// [the Manifold API docs for POST /v0/comment]: https://docs.manifold.markets/api#post-v0comment
func (mc *Client) PostComment(marketId string, pcr PostCommentRequest) error {
	jsonBody, err := json.Marshal(pcr)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	resp, err := mc.makeRequest(context.TODO(), http.MethodPost, requestURL(
		mc.url, postComment,
		marketId,
		""), bodyReader)
	if err != nil {
		return fmt.Errorf("client: error making http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("posting comment failed with status %d", resp.StatusCode)
	}

	return nil
}

// parseResponse takes an HTTP response and a type and attempts to unmarshal
// the body from the response into the given type.
func parseResponse[S any](r *http.Response, s S) (*S, error) {
	if r == nil {
		panic("cannot parse response on nil *http.Response")
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if r.StatusCode != http.StatusOK {
		if strings.Contains(r.Header.Get("Content-Type"), "html") {
			body = body[:100]
		}
		return nil, fmt.Errorf("non-200 status code found: %v, message: %v", r.StatusCode, string(body))
	}

	if len(body) > 0 {
		if err := json.Unmarshal(body, &s); err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
		}
	}

	return &s, nil
}

const Version = "0.2"

func (mc *Client) postRequest(req *http.Request) (*http.Response, error) {
	if mc.key == "" {
		return nil, fmt.Errorf("no API key found")
	}

	req.Header.Set("User-Agent", fmt.Sprintf("github.com/kevinburke/mango/%s go/%s", Version, runtime.Version()))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Key "+mc.key)

	return mc.client.Do(req)
}
