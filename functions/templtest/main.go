package main

import (
	"bytes"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer
	component := hello("John")
	component.Render(context.Background(), &buf)
	output := buf.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       output,
	}
	return res, nil

}

func main() {
	lambda.Start(handler)
}
