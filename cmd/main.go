package main

import (
	"UFG/aws"
	"UFG/types"
	"strconv"
)

func main() {
	pvd := &aws.AwsProvider{}
	pvd.Initialize(makeRequest)
}

func makeRequest(input types.RequestLoop) ([]string, error) {
	arrayName := []string{}
	for loop := 0; loop < input.Input; loop++ {
		arrayName = append(arrayName, "names" + strconv.Itoa(loop))
	}

	return arrayName, nil
}