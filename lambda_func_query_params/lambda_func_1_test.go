package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

// Test with AWS signature
// How to send POST request with json body
// https://stackoverflow.com/questions/57780438/unable-to-set-post-body-in-a-http-request
func Test_with_aws_signature(t *testing.T) {
	creds := credentials.NewStaticCredentials("ACCESS_KEY", "SECRET_KEY", "")
	signer := v4.NewSigner(creds)
	//
	request, body := buildRequest("execute-api", "us-east-1", "")
	fmt.Println(body)
	//
	signer.Sign(request, body, "execute-api", "us-east-1", time.Now())
	//
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Errorf("expect not no error, got %v", err)
	}
	//
	if e, a := http.StatusOK, resp.StatusCode; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	// Close response
	defer resp.Body.Close()
	//
	var respBody ResponseBody
	err2 := json.NewDecoder(resp.Body).Decode(&respBody)
	if err2 != nil {
		t.Errorf("error decoding response body %v", err2)
	}
	//
	fmt.Println(respBody.Message)

	//Body content---------------------------------------------------

	wantBody := "Hello Juan Diaz"
	gotBody := respBody.Message

	if gotBody != wantBody {
		message := "Not as expected Body content. "
		message += fmt.Sprintf("got = %v, want = %v", gotBody, wantBody)

		t.Fatal(message)
		log.Fatal(message)
		t.Errorf("got = %q, want = %q", gotBody, wantBody)
	}
}

func buildRequest(serviceName, region string, body string) (*http.Request, io.ReadSeeker) {

	// Create a new instance of Person
	person := Person{
		FirstName: "Juan",
		LastName:  "Diaz",
	}

	url_query_params := fmt.Sprintf(url_address_lambda_func_query_params+"?firstname=%v&lastname=%v", person.FirstName, person.LastName)

	endpoint := url_query_params
	request, _ := http.NewRequest("POST", endpoint, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")

	return request, strings.NewReader(body)
}
