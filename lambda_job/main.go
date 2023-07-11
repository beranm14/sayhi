package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	countOfCats, _ := strconv.Atoi(os.Getenv("COUNT_OF_CATS"))
	catOutput := ""
	for i := 0; i < countOfCats; i++ {
		fmt.Println("🐈")
		catOutput += "🐈"
	}
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("\"Hello from %s\"", catOutput),
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
