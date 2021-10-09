package api

import (
	"github.com/graphql-go/graphql"
)

var IpAddressType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IpAddress",
		Fields: graphql.Fields{
			"query":       &graphql.Field{Type: graphql.String},
			"status":      &graphql.Field{Type: graphql.String},
			"country":     &graphql.Field{Type: graphql.String},
			"countryCode": &graphql.Field{Type: graphql.String},
			"region":      &graphql.Field{Type: graphql.String},
			"regionName":  &graphql.Field{Type: graphql.String},
			"city":        &graphql.Field{Type: graphql.String},
			"zip":         &graphql.Field{Type: graphql.String},
			"lat":         &graphql.Field{Type: graphql.Float},
			"lon":         &graphql.Field{Type: graphql.Float},
			"timezone":    &graphql.Field{Type: graphql.String},
			"isp":         &graphql.Field{Type: graphql.String},
			"org":         &graphql.Field{Type: graphql.String},
			"as":          &graphql.Field{Type: graphql.String},
			"capital":     &graphql.Field{Type: graphql.String},
			"currency":    &graphql.Field{Type: graphql.String},
			"rates": &graphql.Field{
				Type:    graphql.NewList(RateType),
				Resolve: GetCurrencyRates,
				Args: graphql.FieldConfigArgument{
					"currencies": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
				},
			},
			"phonePrefix": &graphql.Field{Type: graphql.String},
		},
	},
)

var RateType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Rate",
		Fields: graphql.Fields{
			"currencyName": &graphql.Field{Type: graphql.String},
			"rate":         &graphql.Field{Type: graphql.Float},
		},
	},
)

func CreateGraphqlSchema() (graphql.Schema, error) {
	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getIpAddress": &graphql.Field{
					Type:    IpAddressType,
					Resolve: GetIpAddress,
					Args: graphql.FieldConfigArgument{
						"lang": &graphql.ArgumentConfig{Type: graphql.String},
						"ip":   &graphql.ArgumentConfig{Type: graphql.String},
					},
				},
			},
		},
	)

	return graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})
}
