//go:build ignore
// +build ignore

package mango

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: extract non-deterministic tests out as e2e tests, replace them with deterministic unit tests
func TestGetAuthenticatedUser(t *testing.T) {
	expected := User{
		ID:            "igi2zGXsfxYPgB0DJTXVJVmwCOr2",
		CreatedTime:   1639011767273,
		Name:          "Austin",
		Username:      "Austin",
		Url:           "https://manifold.markets/Austin",
		AvatarUrl:     "https://lh3.googleusercontent.com/a-/AOh14GiZyl1lBehuBMGyJYJhZd-N-mstaUtgE4xdI22lLw=s96-c",
		BannerUrl:     "https://images.unsplash.com/photo-1501523460185-2aa5d2a0f981?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1531&q=80",
		Balance:       10000.0,
		TotalDeposits: 10000.0,
		ProfitCached:  ProfitCached{},
		Bio:           "I build Manifold! Always happy to chat; reach out on Discord or find a time on https://calendly.com/austinchen/manifold!",
		Website:       "https://blog.austn.io",
		TwitterHandle: "akrolsmir",
		DiscordHandle: "akrolsmir#4125",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	result, err := mc.GetAuthenticatedUser(context.TODO())
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	b, s := equalUsers(*result, expected)
	if !b {
		t.Errorf(s)
		t.Fail()
	}
}

func TestGetBets(t *testing.T) {
	mc := defaultClientInstance()

	var tests = []struct {
		ui, un, ci, cs, b string
		l                 int64
	}{
		{"", "", "", "", "", defaultLimit},
		{"xN67Q0mAhddL0X9wVYP2YfOrYH42", "", "", "", "", 10},
		{"", "jonny", "", "", "", 10},
		{"", "", "5BOGaVlxLaZt6sdPSUkn", "", "", 10},
		{"", "", "", "dan-stock-permanent", "", 10},
		{"", "", "", "", "6LvbmW25hvrAsjrme209", 10},
		{"", "", "", "", "", 10},
		{"", "", "5BOGaVlxLaZt6sdPSUkn", "dan-stock-permanent", "", 10},
		{"xN67Q0mAhddL0X9wVYP2YfOrYH42", "jonny", "", "", "DgDQOsc7x2KIWTmojZbx", 10},
	}

	for _, test := range tests {
		gbr := GetBetsRequest{
			test.ui, test.un, test.ci, test.cs, test.b, test.l,
		}
		actual, err := mc.GetBets(gbr)
		if err != nil {
			t.Errorf("error getting bets: %v", err)
			t.Fail()
		}
		if int64(len(*actual)) != test.l {
			t.Errorf("incorrect number of bets retrieved")
			t.Fail()
		}
	}
}

func TestGetComments(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		ci, cs string
	}{
		{"5BOGaVlxLaZt6sdPSUkn", ""},
		{"", "dan-stock-permanent"},
		{"5BOGaVlxLaZt6sdPSUkn", "dan-stock-permanent"},
	}

	for _, test := range tests {
		actual, err := mc.GetComments(GetCommentsRequest{test.ci, test.cs})
		if err != nil {
			t.Errorf("error getting comments: %v", err)
			t.Fail()
		}
		if len(*actual) < 1 {
			t.Errorf("incorrect number of comments retrieved")
			t.Fail()
		}
	}
}

func TestGetGroupByID(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		s string
	}{
		{"IlzY3moWwOcpsVZXCVej"},
	}

	for _, test := range tests {
		actual, err := mc.GetGroupById(test.s)
		if err != nil {
			t.Errorf("error getting group by id: %v", err)
			t.Fail()
		}
		if actual.TotalMembers < 1 {
			t.Errorf("incorrect number of members on retrieved group")
			t.Fail()
		}
	}
}

func TestGetGroupBySlug(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		s string
	}{
		{"technology-default"},
	}

	for _, test := range tests {
		actual, err := mc.GetGroupBySlug(test.s)
		if err != nil {
			t.Errorf("error getting group by slug: %v", err)
			t.Fail()
		}
		if actual.TotalMembers < 1 {
			t.Errorf("incorrect number of members on retrieved group")
			t.Fail()
		}
	}
}

func TestGetGroups(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		ui string
	}{
		{"xN67Q0mAhddL0X9wVYP2YfOrYH42"},
		{""},
	}

	for _, test := range tests {
		actual, err := mc.GetGroups(&test.ui)
		if err != nil {
			t.Errorf("error getting groups: %v", err)
			t.Fail()
		}
		if len(*actual) < 1 {
			t.Errorf("incorrect number of groups retrieved")
			t.Fail()
		}
	}
}

func TestGetMarketByID(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		mi string
	}{
		{"5BOGaVlxLaZt6sdPSUkn"},
	}

	for _, test := range tests {
		actual, err := mc.GetMarketByID(context.TODO(), test.mi)
		if err != nil {
			t.Errorf("error getting market by id: %v", err)
			t.Fail()
		}
		if actual.Volume < 1 {
			t.Errorf("incorrect volume on retrieved market")
			t.Fail()
		}
	}
}

func TestGetMarketBySlug(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		ms string
	}{
		{"dan-stock-permanent"},
	}

	for _, test := range tests {
		actual, err := mc.GetMarketBySlug(test.ms)
		if err != nil {
			t.Errorf("error getting market by slug: %v", err)
			t.Fail()
		}
		if actual.Volume < 1 {
			t.Errorf("incorrect volume on retrieved market")
			t.Fail()
		}
	}
}

func TestGetMarkets(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		b string
		l int
	}{
		{"4QTb4cANeQzXNQS9lZnn", 10},
		{"4QTb4cANeQzXNQS9lZnn", defaultLimit},
		{"", defaultLimit},
	}

	for _, test := range tests {
		actual, err := mc.GetMarkets(context.TODO(), GetMarketsRequest{test.b, int64(test.l)})
		if err != nil {
			t.Errorf("error getting markets: %v", err)
			t.Fail()
		}
		if len(*actual) != test.l {
			t.Errorf("incorrect number of markets retrieved")
			t.Fail()
		}
	}
}

func TestGetMarketsForGroup(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		gi string
	}{
		{"IlzY3moWwOcpsVZXCVej"},
	}

	for _, test := range tests {
		actual, err := mc.GetMarketsForGroup(test.gi)
		if err != nil {
			t.Errorf("error getting markets for group: %v", err)
			t.Fail()
		}
		if len(*actual) < 1 {
			t.Errorf("incorrect number of markets for group")
			t.Fail()
		}
	}
}

func TestGetUserByID(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		ui string
	}{
		{"xN67Q0mAhddL0X9wVYP2YfOrYH42"},
	}

	for _, test := range tests {
		actual, err := mc.GetUserByID(test.ui)
		if err != nil {
			t.Errorf("error getting user by id: %v", err)
			t.Fail()
		}
		if actual.Balance < 1 {
			t.Errorf("incorrect balance on retrieved user")
			t.Fail()
		}
	}
}

func TestGetUserByUsername(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		un string
	}{
		{"jonny"},
	}

	for _, test := range tests {
		actual, err := mc.GetUserByUsername(test.un)
		if err != nil {
			t.Errorf("error getting user by username: %v", err)
			t.Fail()
		}
		if actual.Balance < 1 {
			t.Errorf("incorrect balance on retrieved user")
			t.Fail()
		}
	}
}

func defaultClientInstance() *Client {
	return Client{}
}

func TestGetUsers(t *testing.T) {
	mc := defaultClientInstance()
	defer mc.Destroy()

	var tests = []struct {
		b string
		l int
	}{
		{"xN67Q0mAhddL0X9wVYP2YfOrYH42", 10},
		{"xN67Q0mAhddL0X9wVYP2YfOrYH42", defaultLimit},
		{"", defaultLimit},
	}

	for _, test := range tests {
		actual, err := mc.GetUsers(GetUsersRequest{test.b, int64(test.l)})
		if err != nil {
			t.Errorf("error getting users: %v", err)
			t.Fail()
		}
		if len(*actual) != test.l {
			t.Errorf("incorrect number of users retrieved")
			t.Fail()
		}
	}
}

func TestGetMarketPositions(t *testing.T) {
	expected := []ContractMetric{
		{
			ContractId: "1",
			From: map[string]Period{
				"2022-01-01": {
					Profit:        100.0,
					ProfitPercent: 10.0,
					Invested:      1000.0,
					PrevValue:     1000.0,
					Value:         1100.0,
				},
			},
			HasNoShares:   false,
			HasShares:     true,
			HasYesShares:  false,
			Invested:      1000.0,
			Loan:          0.0,
			MaxShares:     "",
			Payout:        100.0,
			Profit:        100.0,
			ProfitPercent: 10.0,
			TotalShares: map[string]float64{
				"NO":  100.0,
				"YES": 0.0,
			},
			UserId:        "user1",
			UserName:      "John Doe",
			UserUsername:  "johndoe",
			UserAvatarUrl: "https://example.com/avatar.png",
			LastBetTime:   1641004800,
		},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	result, err := mc.GetMarketPositions(context.TODO(), GetMarketPositionsRequest{"1", "desc", 0, 100, "user1"})
	if err != nil {
		t.Errorf("error getting market positions: %v", err)
		t.Fail()
	}

	if len(*result) != len(expected) {
		t.Errorf("unexpected result length: got %d, want %d", len(*result), len(expected))
	}

	for i, contract := range *result {
		b, s := equalContractMetrics(contract, expected[i])
		if !b {
			t.Errorf(s)
		}
	}
}

func TestSearchMarkets(t *testing.T) {
	expected := []FullMarket{
		{
			Id:                    "123",
			CreatorId:             "456",
			CreatorUsername:       "jonny",
			CreatorName:           "Jonny",
			CreatedTime:           0,
			CreatorAvatarUrl:      "https://127.0.0.1",
			CloseTime:             1,
			Question:              "How much wood would a woodchuck chuck of a woodchuck could chuck wood?",
			Answers:               nil,
			Tags:                  nil,
			Url:                   "",
			Pool:                  Pool{},
			Probability:           50,
			P:                     0,
			TotalLiquidity:        0,
			OutcomeType:           Binary,
			Mechanism:             "dpm-2",
			Volume:                10000,
			Volume24Hours:         10000,
			IsResolved:            false,
			Resolution:            "",
			ResolutionTime:        0,
			ResolutionProbability: 0,
			LastUpdatedTime:       2,
			TextDescription:       "Will resolve based on some totally arbitrary criteria I pick at resolution time",
		},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	result, err := mc.SearchMarkets("apple", "banana", "celery", "damson")
	if err != nil {
		t.Errorf("error searching markets: %v", err)
		t.Fail()
	}

	if len(*result) != len(expected) {
		t.Errorf("unexpected result length: got %d, want %d", len(*result), len(expected))
	}

	for i, market := range *result {
		b, s := equalFullMarkets(market, expected[i])
		if !b {
			t.Errorf(s)
		}
	}
}

func TestPostBet(t *testing.T) {
	pbr := PostBetRequest{
		Amount:     10,
		ContractId: "abc123",
		Outcome:    "YES",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	_, err := mc.PostBet(context.Background(), pbr)
	if err != nil {
		t.Errorf("error posting bet: %v", err)
		t.Fail()
	}
}

func TestCancelBet(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	err := mc.CancelBet(context.Background(), "123abc")
	if err != nil {
		t.Errorf("error cancelling bet: %v", err)
		t.Fail()
	}
}

func TestCreateMarket(t *testing.T) {
	pmr := PostMarketRequest{
		OutcomeType:         Binary,
		Question:            "How much wood would a woodchuck chuck of a woodchuck could chuck wood?",
		Description:         "Will resolve based on some totally arbitrary criteria I pick at resolution time",
		DescriptionHtml:     "",
		DescriptionMarkdown: "",
		CloseTime:           1,
		Visibility:          "",
		GroupId:             "",
		InitialProb:         1,
		Min:                 0,
		Max:                 10,
		IsLogScale:          false,
		InitialVal:          0,
		Answers:             nil,
	}

	expected := marketIdResponse{Id: "123marketId"}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	resp, err := mc.CreateMarket(context.TODO(), pmr)
	if err != nil {
		t.Errorf("error creating market: %v", err)
		t.Fail()
	}

	if *resp != expected.Id {
		t.Errorf("market ID responses don't match, got: %v expected: %v", *resp, expected)
		t.Fail()
	}
}

func TestAddLiquidity(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	err := mc.AddLiquidity("123marketid", 9999)
	if err != nil {
		t.Errorf("error adding liquidity: %v", err)
		t.Fail()
	}
}

func TestCloseMarket(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	err := mc.CloseMarket("123marketid", nil)
	if err != nil {
		t.Errorf("error closing market: %v", err)
		t.Fail()
	}
}

func TestAddMarketToGroup(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	err := mc.AddMarketToGroup(context.TODO(), "123marketid", "123groupid")
	if err != nil {
		t.Errorf("error adding market to group: %v", err)
		t.Fail()
	}
}

func TestResolveMarket(t *testing.T) {
	rmr := ResolveMarketRequest{
		Outcome:        "YES",
		Resolutions:    nil,
		ProbabilityInt: 0,
		Value:          0,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	err := mc.ResolveMarket("123marketid", rmr)
	if err != nil {
		t.Errorf("error resolving market: %v", err)
		t.Fail()
	}
}

func TestSellShares(t *testing.T) {
	ssr := SellSharesRequest{
		Outcome: "YES",
		Shares:  10000,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	err := mc.SellShares("123marketid", ssr)
	if err != nil {
		t.Errorf("error selling shares: %v", err)
		t.Fail()
	}
}

func TestPostComment(t *testing.T) {
	pcr := PostCommentRequest{
		ContractId: "123contractid",
		Content:    "insert snarky comment here",
		Html:       "",
		Markdown:   "",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	mc := ClientInstance(server.Client(), &server.URL, nil)
	defer mc.Destroy()

	err := mc.PostComment("123marketid", pcr)
	if err != nil {
		t.Errorf("error posting comment: %v", err)
		t.Fail()
	}
}
