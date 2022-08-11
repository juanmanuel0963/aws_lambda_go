package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var (
	DefaultHTTPAddress = "https://8syalbja7g.execute-api.us-east-1.amazonaws.com/hello_world_user"
)

//To review
//https://mailazy.com/blog/http-request-golang-with-best-practices/

func Test_hello_world_user(t *testing.T) {

	t.Run("POST_request", func(t *testing.T) {

		// Create a new instance of Person
		person := Person{
			FirstName: "Juan",
			LastName:  "Diaz",
		}

		// Marshal it into JSON prior to requesting
		personJSON, err := json.Marshal(person)

		//POST request---------------------------------------------------

		// Make request with marshalled JSON as the POST body
		resp, err := http.Post(DefaultHTTPAddress, "application/json", bytes.NewBuffer(personJSON))

		if err != nil {
			message := "Unsuccessfull POST request to " + DefaultHTTPAddress
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

		//Unmarchalling response body---------------------------------------------------

		body, err := ioutil.ReadAll(resp.Body)

		var result ResponseBody
		err = json.Unmarshal([]byte(body), &result)

		if err != nil {
			message := "Error unmarshaling body content from request."
			t.Fatal(message)
			log.Fatal(message)
			t.Errorf(message)
		}

		//Body content---------------------------------------------------

		wantBody := "Hello Juan Diaz"
		gotBody := result.Message

		if gotBody != wantBody {
			message := "Not as expected Body content. "
			message += fmt.Sprintf("got = %v, want = %v", gotBody, wantBody)

			t.Fatal(message)
			log.Fatal(message)
			t.Errorf("got = %q, want = %q", gotBody, wantBody)
		}

	})
}
