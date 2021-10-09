package api

import (
	"log"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGraphql(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		Body:    `{"query":"query{getIpAddress(lang:\"es\"){city country query currency rates(currencies:[\"CZK\", \"EUR\"]){currencyName rate}}}","variables":{}}`,
		Headers: map[string]string{"X-Forwarded-For": "37.188.168.186"},
	}

	res, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.Body != `{"data":{"getIpAddress":{"city":"Prague","country":"Czechia","currency":"CZK","query":"37.188.168.186","rates":[{"currencyName":"CZK","rate":1},{"currencyName":"EUR","rate":0.038683917}]}}}` {
		log.Print(res.Body)
		log.Fatal("body is not equal")
	}
}
