package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

var (
	SERVICE_URL = "https://3vefaztvjh.execute-api.us-east-1.amazonaws.com/lambda_func"
	ACCESS_KEY  = ""
	SECRET_KEY  = ""
)

// Test without AWS Signature
func Test_without_aws_signature(t *testing.T) {

	//https://docs.aws.amazon.com/general/latest/gr/sigv4-signed-request-examples.html#sig-v4-examples-post
	//https://github.com/aws/aws-sdk-go/tree/main/aws/signer/v4

	t.Run("GET_request_auth", func(t *testing.T) {

		//GET request---------------------------------------------------

		client := http.Client{}
		req, err := http.NewRequest("GET", SERVICE_URL, nil)

		if err != nil {
			message := "Unsuccessfull GET request to " + SERVICE_URL
			t.Fatal(message)
			log.Fatal(message)
			t.Errorf(message)
		}

		resp, err := client.Do(req)

		if err != nil {
			message := "Unsuccessfull GET request to " + SERVICE_URL
			t.Fatal(message)
			log.Fatal(message)
			t.Errorf(message)
		}

		//Response code 200-----------------------------------------------

		wantCode := 200
		gotCode := resp.StatusCode

		if gotCode != wantCode {
			message := "Response code 200 expected. "
			message += fmt.Sprintf("Got = %v, Want = %v", gotCode, wantCode)

			t.Fatal(message)
			log.Fatal(message)
			t.Errorf("Got = %q, Want = %q", gotCode, wantCode)
		}
	})

}
