package api

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"

	"github.com/joho/godotenv"
	"github.com/vladimir-chernenko/fixerapi"
)

var fc *fixerapi.FixerClient
var schema graphql.Schema

func init() {
	err := godotenv.Load()

	apiKey, ok := os.LookupEnv("FIXER_API_KEY")

	if !ok {
		log.Fatal("FIXER_API_KEY must be set for testing")
	}

	fc = fixerapi.NewFixerClient(apiKey)

	schema, err = CreateGraphqlSchema()

	if err != nil {
		log.Fatal(err)
	}
}

type GraphqlQuery struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	q := GraphqlQuery{}

	err := json.Unmarshal([]byte(req.Body), &q)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 400}, err
	}

	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  q.Query,
		VariableValues: q.Variables,
		OperationName:  q.OperationName,
		Context:        context.WithValue(context.Background(), "userIPs", req.Headers["X-Forwarded-For"]),
	})

	response, err := json.Marshal(result)

	res := events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}

	return res, err
}

func main() {
	lambda.Start(Handler)
}
