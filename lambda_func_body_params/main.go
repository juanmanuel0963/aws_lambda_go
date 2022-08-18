package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler_lambda_func_body_params)
}

func Handler_lambda_func_body_params(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("Hello world user")

	log.Println("Request: ", request)

	log.Println("Request body1: ", request.Body)

	log.Println("Request body2: ", []byte(request.Body))

	var person Person

	err := json.Unmarshal([]byte(request.Body), &person)

	if err != nil {

		log.Println("Request body3: ", err)

		return events.APIGatewayProxyResponse{}, err
	}

	log.Println("Request body3: ", person.FirstName, person.LastName)

	msg := fmt.Sprintf("Hello %v %v ", person.FirstName, person.LastName)

	responseBody := ResponseBody{
		Message: msg,
	}

	jbytes, err := json.Marshal((responseBody))

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jbytes),
	}

	return response, nil
}

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type ResponseBody struct {
	Message string `json:"message"`
}
