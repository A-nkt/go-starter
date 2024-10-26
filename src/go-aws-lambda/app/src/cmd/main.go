package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name     string `json:"name"`
	Describe bool   `json:"describe"`
}

func handler(ctx context.Context, event MyEvent) {
	fmt.Println("Hello, World!")
}

func main() {

	lambda.Start(handler)
}
