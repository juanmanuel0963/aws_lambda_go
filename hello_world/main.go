package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//log.Println("Hello world")
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//log.Println("Hello world")
	var person Person

	err := json.Unmarshal([]byte(request.Body), &person)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	msg := fmt.Sprintf("Hello %v %v", *person.FirstName, *person.LastName)

	responseBody := ResponseBody{
		Message: &msg,
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
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type ResponseBody struct {
	Message *string `json:"message"`
}
