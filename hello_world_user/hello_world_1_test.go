package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

//Test with AWS signature
//How to send POST request with json body
//https://stackoverflow.com/questions/57780438/unable-to-set-post-body-in-a-http-request
func TestSignWithRequestBodyTest(t *testing.T) {
	creds := credentials.NewStaticCredentials("AKIA4A7LNDSX2Y7HIECJ", "bqF9o7EFhvhk9cvPkG8P1YOfF9hGa8E8LdgLedH5", "")
	signer := v4.NewSigner(creds)

	// Create a new instance of Person
	person := Person{
		FirstName: "Juan Manuel",
		LastName:  "Diaz Ortiz",
	}
	// Marshal it into JSON prior to requesting
	jsonBody, err := json.Marshal(person)
	fmt.Println(string(jsonBody))
	//
	request, body := buildRequest("execute-api", "us-east-1", string(jsonBody))
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
}

func buildRequest(serviceName, region string, body string) (*http.Request, io.ReadSeeker) {

	endpoint := DefaultHTTPAddress_user
	request, _ := http.NewRequest("POST", endpoint, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")

	return request, strings.NewReader(body)
}
