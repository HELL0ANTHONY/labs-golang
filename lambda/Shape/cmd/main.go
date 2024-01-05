package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/handler"
)

func main() {
	h := handler.New()
	lambda.Start(h.Handle)
}
