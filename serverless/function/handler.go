package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func handler(ctx context.Context) (Response, error) {
	return Response{
		StatusCode: 200,
		Body:       time.Now().Format(time.RFC3339),
	}, nil
}

func main() {
	lambda.Start(handler)
}
