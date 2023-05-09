package mango

var betResponse = []byte(`
{
    "orderAmount": 2,
    "amount": 2,
    "shares": 2.223856205901484,
    "limitProb": 0.1,
    "isFilled": true,
    "isCancelled": false,
    "fills": [
        {
            "matchedBetId": null,
            "shares": 2.223856205901484,
            "amount": 2,
            "timestamp": 1683604157926
        }
    ],
    "contractId": "ZIOKo6DMQszoaQgTOtda",
    "outcome": "NO",
    "probBefore": 0.1008648208171883,
    "probAfter": 0.10045820256853283,
    "loanAmount": 0,
    "createdTime": 1683604157926,
    "fees": {
        "creatorFee": 0,
        "platformFee": 0,
        "liquidityFee": 0
    },
    "isAnte": false,
    "isRedemption": false,
    "isChallenge": false,
    "visibility": "public",
    "betId": "FEBRIVsLWHNb9EMmw2nY"
}
`)

var marketResponse = []byte(`
{
    "id": "3n13Itp3fELqN27x8DRb",
    "creatorId": "VdbwzTgpBYV2lY2aw5T6m4ISbAD3",
    "creatorUsername": "prognostic8r",
    "creatorName": "prognostic8r",
    "createdTime": 1675941658643,
    "creatorAvatarUrl": "https://firebasestorage.googleapis.com/v0/b/mantic-markets.appspot.com/o/user-images%2Fsports%2F-MgwhwW5au.png?alt=media&token=b14df256-f9db-4969-b9e3-31dcacb243b6",
    "closeTime": 1687129200000,
    "question": "Will the Los Angeles Lakers win the finals in the 2022-2023 NBA season?",
    "tags": [],
    "url": "https://manifold.markets/prognostic8r/will-the-los-angeles-lakers-win-the",
    "pool": {
        "NO": 339.11890952039124,
        "YES": 239.0064425589963
    },
    "probability": 0.13311080465201314,
    "p": 0.09765204809505523,
    "totalLiquidity": 370,
    "outcomeType": "BINARY",
    "mechanism": "cpmm-1",
    "volume": 3829.751624852715,
    "volume24Hours": 2,
    "isResolved": false,
    "lastUpdatedTime": 1683605432101,
    "description": {
        "type": "doc",
        "content": [
            {
                "type": "paragraph",
                "content": [
                    {
                        "text": "This market will resolve 'YES' if the Los Angeles Lakers win the NBA Finals in the 2022-2023 season. Starting odds are taken from ",
                        "type": "text"
                    },
                    {
                        "text": "FiveThirtyEight's forecast",
                        "type": "text",
                        "marks": [
                            {
                                "type": "link",
                                "attrs": {
                                    "href": "https://projects.fivethirtyeight.com/2023-nba-predictions/",
                                    "class": null,
                                    "target": "_blank"
                                }
                            }
                        ]
                    },
                    {
                        "text": " on the date of this market's creation.",
                        "type": "text"
                    }
                ]
            },
            {
                "type": "paragraph",
                "content": [
                    {
                        "text": "Market created with ",
                        "type": "text"
                    },
                    {
                        "text": "manifoldr",
                        "type": "text",
                        "marks": [
                            {
                                "type": "link",
                                "attrs": {
                                    "href": "https://github.com/jcblsn/manifoldr",
                                    "class": null,
                                    "target": "_blank"
                                }
                            }
                        ]
                    },
                    {
                        "text": ".",
                        "type": "text"
                    }
                ]
            }
        ]
    },
    "coverImageUrl": "https://firebasestorage.googleapis.com/v0/b/mantic-markets.appspot.com/o/dream%2FdNg9KTjJim.png?alt=media&token=a72aa58b-5005-41ea-ac7a-31263938c2a6",
    "textDescription": "This market will resolve 'YES' if the Los Angeles Lakers win the NBA Finals in the 2022-2023 season. Starting odds are taken from FiveThirtyEight's forecast on the date of this market's creation.\n\nMarket created with manifoldr."
}
`)
