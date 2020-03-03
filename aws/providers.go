package aws

import (
	"UFG/types"
	"github.com/aws/aws-lambda-go/lambda"
)

type AwsProvider struct {
	handlerFunc func(input types.RequestLoop) ([]string, error)
	Input     	*types.RequestLoop
}

func (a *AwsProvider) Initialize(handler func(input types.RequestLoop) ([]string, error)) {
	a.handlerFunc = handler
	lambda.Start(a.makeRequest)
}

func (a *AwsProvider) makeRequest(input types.RequestLoop) ([]string, error) {
	a.Input = &input
	return a.handlerFunc(*a.Input)
}