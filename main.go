package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type HelloLambdaResponse struct {
	Greeting    string    `json:"greeting"`
	CurrentTime time.Time `json:"current_time"`
}

func handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := &HelloLambdaResponse{
		Greeting:    "Hello World!",
		CurrentTime: time.Now().UTC(),
	}

	responseBody, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil

}

func main() {
	lambda.Start(handle)
}
