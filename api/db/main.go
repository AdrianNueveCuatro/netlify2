package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Stores a handle to the collection being used by the Lambda function
type Connection struct {
	collection *mongo.Collection
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            "DB Hello, World!",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://dbUser:87rt45th67>@cluster0.swvad.mongodb.net/Books?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	lambda.Start(handler)
}
