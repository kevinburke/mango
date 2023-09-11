package mango

import "fmt"

// ProfitCached represents the profit for a given [User]
type ProfitCached struct {
	Weekly  float64 `json:"weekly"`
	Daily   float64 `json:"daily"`
	AllTime float64 `json:"allTime"`
	Monthly float64 `json:"monthly"`
}

// GetUsersRequest represents the optional parameters that can be supplied to
// get users via the API
type GetUsersRequest struct {
	Before string `json:"before,omitempty"`
	Limit  int64  `json:"limit,omitempty"`
}

// User represents a User object in the Manifold backend.
//
// This type isn't documented by Manifold and its structure was inferred from API calls.
type User struct {
	ID            string       `json:"id"`
	CreatedTime   int64        `json:"createdTime"`
	Name          string       `json:"name"`
	Username      string       `json:"username"`
	Url           string       `json:"url"`
	AvatarUrl     string       `json:"avatarUrl"`
	BannerUrl     string       `json:"bannerUrl"`
	Balance       float64      `json:"balance"`
	TotalDeposits float64      `json:"totalDeposits"`
	ProfitCached  ProfitCached `json:"profitCached"`
	Bio           string       `json:"bio,omitempty"`
	Website       string       `json:"website,omitempty"`
	TwitterHandle string       `json:"twitterHandle,omitempty"`
	DiscordHandle string       `json:"discordHandle,omitempty"`
}

func equalUsers(u1, u2 User) (bool, string) {
	if u1.ID != u2.ID {
		return false, fmt.Sprintf("user id is not equal: %v & %v", u1.ID, u2.ID)
	}
	if u1.CreatedTime != u2.CreatedTime {
		return false, fmt.Sprintf("user CreatedTime is not equal: %v & %v", u1.CreatedTime, u2.CreatedTime)
	}
	if u1.Name != u2.Name {
		return false, fmt.Sprintf("user Name is not equal: %v & %v", u1.Name, u2.Name)
	}
	if u1.Username != u2.Username {
		return false, fmt.Sprintf("user Username is not equal: %v & %v", u1.Username, u2.Username)
	}
	if u1.AvatarUrl != u2.AvatarUrl {
		return false, fmt.Sprintf("user AvatarUrl is  not equal: %v & %v", u1.AvatarUrl, u2.AvatarUrl)
	}
	if u1.BannerUrl != u2.BannerUrl {
		return false, fmt.Sprintf("user BannerUrl is not equal: %v & %v", u1.BannerUrl, u2.BannerUrl)
	}
	if u1.Balance != u2.Balance {
		return false, fmt.Sprintf("user Balance is not equal: %v & %v", u1.Balance, u2.Balance)
	}
	if u1.TotalDeposits != u2.TotalDeposits {
		return false, fmt.Sprintf("user TotalDeposits is not equal: %v & %v", u1.TotalDeposits, u2.TotalDeposits)
	}
	if u1.Bio != u2.Bio {
		return false, fmt.Sprintf("user Bio is not equal: %v & %v", u1.Bio, u2.Bio)
	}
	if u1.Website != u2.Website {
		return false, fmt.Sprintf("user Website is not equal: %v & %v", u1.Website, u2.Website)
	}
	if u1.TwitterHandle != u2.TwitterHandle {
		return false, fmt.Sprintf("user TwitterHandle is not equal: %v & %v", u1.TwitterHandle, u2.TwitterHandle)
	}
	if u1.DiscordHandle != u2.DiscordHandle {
		return false, fmt.Sprintf("user DiscordHandle is not equal: %v & %v", u1.DiscordHandle, u2.DiscordHandle)
	}
	if !equalProfitCached(u1.ProfitCached, u2.ProfitCached) {
		return false, fmt.Sprintf("user ProfitCached is not equal: %v & %v", u1.ProfitCached, u2.ProfitCached)
	}

	return true, ""
}

func equalProfitCached(p1, p2 ProfitCached) bool {
	return p1.Weekly == p2.Weekly &&
		p1.Daily == p2.Daily &&
		p1.AllTime == p2.AllTime &&
		p1.Monthly == p2.Monthly
}
