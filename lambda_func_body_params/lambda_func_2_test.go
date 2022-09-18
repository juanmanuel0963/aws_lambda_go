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

func Test_local_server(t *testing.T) {
	//signing credentials
	creds := credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, "")
	signer := v4.NewSigner(creds)

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
		var personGot Person

		err2 := json.Unmarshal(bodyNew, &personGot)
		if err2 != nil {
			t.Errorf("expect no error, got %v", err)
		}
		fmt.Println(personGot)

		//----------------
		r.Body.Close()

		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}

		if e, a := person, personGot; !reflect.DeepEqual(e, a) {
			t.Errorf("expect %v, got %v", e, a)
		}

		//----------------
		w.WriteHeader(http.StatusOK)

		fmt.Println(http.StatusOK)

		//----------------
		msg := fmt.Sprintf("Hello %v %v ", person.FirstName, person.LastName)

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
