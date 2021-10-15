package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/efarrer/page2pod/authentication"
	"github.com/efarrer/page2pod/htmlform"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	formData, err := htmlform.Parse(req.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("Unable to parse form data: %w", err)
	}

	svc := secretsmanager.New(session.New())
	err = authentication.Authenticate(svc, formData.Username, formData.Password)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 401,
			Body:       "Unable to authenticate " + formData.Username,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       "Hello " + formData.Username,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
