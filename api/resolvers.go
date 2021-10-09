package api

import (
	"strings"

	"github.com/vladimir-chernenko/ipapi"

	"github.com/graphql-go/graphql"
)

func GetCurrencyRates(params graphql.ResolveParams) (interface{}, error) {
	symbols := []string{}

	for _, symbol := range params.Args["currencies"].([]interface{}) {
		symbols = append(symbols, symbol.(string))
	}

	return fc.ConvertCurrency(params.Source.(ipapi.IpDetails).Currency, symbols)
}

func GetIpAddress(params graphql.ResolveParams) (interface{}, error) {
	userIPs := params.Context.Value("userIPs").(string)
	ip := strings.Split(userIPs, ",")[0]
	return ipapi.GetIpDetails(ip)
}
