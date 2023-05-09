package mango

type betResponse = []byte(`
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
