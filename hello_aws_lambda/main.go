package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() (string, error) {
	return "Hello", nil
}

func main() {
	lambda.Start(Handler)
}
