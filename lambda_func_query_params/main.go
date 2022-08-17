package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler_lambda_func_query_params)
}

//ctx context.Context,
func Handler_lambda_func_query_params(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//
	log.Println("Into the function")

	var personGot Person
	msg := ""

	//-------firstName
	firstName, found := request.QueryStringParameters["firstname"]

	if !found {
		log.Println("Not found!")
		msg += fmt.Sprintf("firstname: %v ", "Not found")
	} else {
		log.Println("firstname found!")

		// query parameters are typically URL encoded so to get the value
		value, err := url.QueryUnescape(firstName)
		if nil != err {
			log.Println("QueryUnescape error!")
			return events.APIGatewayProxyResponse{}, err
		}

		personGot.FirstName = value
		msg += fmt.Sprintf("Hello %v ", personGot.FirstName)
		log.Println(msg)
	}

	//-------lastName
	lastName, found := request.QueryStringParameters["lastname"]

	if !found {
		log.Println("Not found!")
		msg += fmt.Sprintf("lastname: %v ", "Not found")
	} else {
		log.Println(" lastname found!")

		// query parameters are typically URL encoded so to get the value
		value, err := url.QueryUnescape(lastName)
		if nil != err {
			log.Println("QueryUnescape error!")
			return events.APIGatewayProxyResponse{}, err
		}

		personGot.LastName = value
		msg += fmt.Sprintf("%v", personGot.LastName)
		log.Println(msg)
	}

	//-------responseBody

	responseBody := ResponseBody{
		Message: msg,
	}

	jsonBody, err := json.Marshal((responseBody))
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	log.Println("jsonBody: " + string(jsonBody))

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		//Body:       "Hello there",
		Body: string(jsonBody),
	}

	log.Println("StatusCode: 200")

	return response, nil
}

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type ResponseBody struct {
	Message string `json:"message"`
}
