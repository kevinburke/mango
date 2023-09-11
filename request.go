package mango

import (
	"log"
	"net/url"
)

const base string = "https://manifold.markets/api"
const version string = "v0/"

const defaultLimit = 1000

const getBets string = "bets/"
const getComments string = "comments/"
const getGroupBySlug string = "group/"
const getGroupByID string = "group/by-id/"
const getGroups string = "groups/"
const getMarketBySlug string = "slug/"
const getMarketByID string = "market/"
const getMarkets string = "markets/"
const getSearchMarkets string = "search-markets/"
const getMe string = "me"
const getUserByUsername string = "user/"
const getUserByID string = "user/by-id/"
const getUsers string = "users/"

const postBet string = "bet/"
const postCancellation string = "bet/cancel/"
const postMarket string = "market"
const postComment string = "comment/"

const marketsSuffix = "/markets/"
const positionsSuffix = "/positions"
const liquiditySuffix = "/liquidity/"
const closureSuffix = "/close/"
const groupSuffix = "/group/"
const resolutionSuffix = "/resolve/"
const sellSuffix = "/sell/"

// requestURL returns a fully-formed URL that HTTP requests can be sent to.
// It includes the base domain, path, and any query parameters supplied.
func requestURL(base, path, value, suffix string, params ...string) string {
	// if query parameters are supplied, they must be in key:value pairs
	if len(params)%2 != 0 {
		log.Println("number of params passed to requestURL() must be divisible by 2")
		return ""
	}

	query, err := url.Parse(base + "/" + version)
	if err != nil {
		log.Fatalf("error parsing base URL: %v", err)
	}

	query.Path += path

	if value != "" {
		query.Path += value
	} else {
		ps := url.Values{}

		for i := 0; i < len(params); i += 2 {
			if params[i+1] != "" {
				ps.Add(params[i], params[i+1])
			}
		}

		query.RawQuery = ps.Encode()
	}

	// the suffix represents part of the path that is supplied after an identifier
	// eg for the path: /v0/market/[marketId]/add-liquidity
	// the string "add-liquidity" would be the suffix
	if suffix != "" {
		query.Path += suffix
	}

	return query.String()
}
