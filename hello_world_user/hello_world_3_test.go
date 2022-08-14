package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

func TestSignWithRequestBody(t *testing.T) {
	creds := credentials.NewStaticCredentials("AKIA4A7LNDSX2Y7HIECJ", "bqF9o7EFhvhk9cvPkG8P1YOfF9hGa8E8LdgLedH5", "")
	signer := v4.NewSigner(creds)

	//expectBody := []byte("abc123")
	//fmt.Println(string(expectBody))

	// Create a new instance of Person
	person := Person{
		FirstName: "Juan",
		LastName:  "Diaz",
	}
	// Marshal it into JSON prior to requesting
	jsonBody, err := json.Marshal(person)
	fmt.Println(string(jsonBody))

	//Create the body
	body := strings.NewReader(string(jsonBody))

	//Create the Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//-------------
		bodyNew, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
		fmt.Println(string(bodyNew))

		//--------------
		var person Person

		err2 := json.Unmarshal(bodyNew, &person)
		if err2 != nil {
			t.Errorf("expect no error, got %v", err)
		}
		fmt.Println(person)

		//----------------
		r.Body.Close()

		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}

		if e, a := body, bodyNew; !reflect.DeepEqual(e, a) {
			t.Errorf("expect %v, got %v", e, a)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Println(http.StatusOK)
	}))
	defer server.Close()

	//Create the request
	request, err := http.NewRequest("POST", server.URL, body)

	if err != nil {
		t.Errorf("expect not no error, got %v", err)
	}

	//Signing
	_, err = signer.Sign(request, body, "execute-api", "us-east-1", time.Now())
	if err != nil {
		t.Errorf("expect not no error, got %v", err)
	}

	//Request execution
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Errorf("expect not no error, got %v", err)
	}

	fmt.Println(resp.Body)

	//Comparisson
	if e, a := http.StatusOK, resp.StatusCode; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}
