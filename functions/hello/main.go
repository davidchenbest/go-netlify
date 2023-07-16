package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "",
	}
	switch request.HTTPMethod {
	case "POST":
		// var data MyData
		var data map[string]interface{}
		b := []byte(request.Body)
		err := json.Unmarshal(b, &data)
		if err != nil {
			fmt.Println(err)
			res.Body = "Invalid JSON"
			res.StatusCode = 400

		}
		res.Headers = map[string]string{"Content-Type": "text/plain"}
		res.Body = "This is post" + request.Body
		res.IsBase64Encoded = false

	case "PATCH":
		person := map[string]any{"name": "Alice", "age": 30}
		jsonData, err := json.Marshal(person)
		if err != nil {
			fmt.Println("Error:", err)
			res.Body = "Invalid JSON"
			res.StatusCode = 400
		}
		res.Headers = map[string]string{"Content-Type": "application/json"}
		res.Body = string(jsonData)
		res.IsBase64Encoded = false
	}
	return res, nil

}

func main() {
	lambda.Start(handler)
}
