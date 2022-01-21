package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Body: ", input.Body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World",
	}, nil
}

func main() {
	lambda.Start(Handle)
}
