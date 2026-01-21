package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Response represents the Lambda function response
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Request represents the Lambda function request
type Request struct {
	Name string `json:"name,omitempty"`
}

// handler is the Lambda function handler
func handler(ctx context.Context, request Request) (Response, error) {
	// Log the incoming request
	log.Printf("Processing request: %+v", request)

	// Default name if not provided
	name := "world"
	if request.Name != "" {
		name = request.Name
	}

	// Create response
	message := fmt.Sprintf("Hello, %s!", name)
	response := Response{
		Message: message,
		Status:  200,
	}

	log.Printf("Returning response: %+v", response)
	return response, nil
}

func main() {
	// Start the Lambda handler
	lambda.Start(handler)
}
