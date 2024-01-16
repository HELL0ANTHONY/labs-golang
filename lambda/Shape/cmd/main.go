package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/internal/processor"
	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/internal/repository"
	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/handler"
	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/validation"
)

func main() {
	r := repository.NewDynamoDBRepository()
	p := processor.New(r)
	v := validation.New()
	h := handler.New(p, v)

	lambda.Start(h.Handle)
}
