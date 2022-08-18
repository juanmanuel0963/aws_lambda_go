package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

func Test_local_server(t *testing.T) {
	//signing credentials
	creds := credentials.NewStaticCredentials("AKIA4A7LNDSX2Y7HIECJ", "bqF9o7EFhvhk9cvPkG8P1YOfF9hGa8E8LdgLedH5", "")
	signer := v4.NewSigner(creds)

	// Create a new instance of Person
	person := Person{
		FirstName: "Juan",
		LastName:  "Diaz",
	}

	//Create the body
	body := strings.NewReader("")

	//Declare the srv variable before using the variable.
	var server *httptest.Server

	//Create the Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var personGot Person
		msg := ""

		QueryStringParameters, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			fmt.Fprintf(w, "Invalid query string parameters")
			t.Errorf("Invalid query string parameters Err: %s", err)
		}

		fmt.Println(QueryStringParameters)

		//First name
		firstname := QueryStringParameters.Get("firstname")

		if len(firstname) == 0 {
			fmt.Fprintf(w, "Missing firstname")
			t.Errorf("Missing firstname")
		}

		fmt.Println(firstname)

		//Last name
		lastname := QueryStringParameters.Get("lastname")

		if len(lastname) == 0 {
			fmt.Fprintf(w, "Missing lastname")
			t.Errorf("Missing lastname")
		}

		fmt.Println(lastname)

		//-------------

		personGot.FirstName = firstname
		personGot.LastName = lastname

		//-------------

		if e, a := person, personGot; !reflect.DeepEqual(e, a) {
			t.Errorf("expect %v, got %v", e, a)
		}

		//----------------
		w.WriteHeader(http.StatusOK)
		fmt.Println(http.StatusOK)

		//----------------
		msg = fmt.Sprintf("Hello %v %v ", person.FirstName, person.LastName)

		responseBody := ResponseBody{
			Message: msg,
		}

		jsonBody, err := json.Marshal((responseBody))
		if err != nil {
			t.Errorf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonBody)
		fmt.Println(jsonBody)

		return
	}))
	defer server.Close()

	local_server_url_query_params := fmt.Sprintf(server.URL+"?firstname=%v&lastname=%v", person.FirstName, person.LastName)

	//Create the request
	request, err := http.NewRequest("POST", local_server_url_query_params, body)

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

	//Comparisson
	if e, a := http.StatusOK, resp.StatusCode; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	fmt.Println(resp.Body)

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
