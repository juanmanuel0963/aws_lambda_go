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
	DefaultHTTPGetAddress = "https://8syalbja7g.execute-api.us-east-1.amazonaws.com/hello_world_user"
)

//To review
//https://mailazy.com/blog/http-request-golang-with-best-practices/

func Test_hello_world_user(t *testing.T) {

	t.Run("Invalid URL", func(t *testing.T) {

		_, err := http.Get(DefaultHTTPGetAddress)

		if err != nil {
			t.Fatal("Invalid URL")
			log.Fatal("Invalid URL")
			t.Errorf("Invalid URL")
		}
	})

	t.Run("Non 200 response", func(t *testing.T) {

		// Create a new instance of Person
		person := Person{
			FirstName: "Juan",
			LastName:  "Diaz",
		}

		// Marshal it into JSON prior to requesting
		personJSON, err := json.Marshal(person)

		// Make request with marshalled JSON as the POST body
		resp, err := http.Post(DefaultHTTPGetAddress, "application/json",
			bytes.NewBuffer(personJSON))

		if err != nil {
			t.Error("Could not make POST request to http")
		}

		want := 200
		got := resp.StatusCode

		if got != want {

			res := fmt.Sprintf("got = %v, want = %v", got, want)

			t.Fatal(res)
			log.Fatal(res)
			t.Errorf("got = %q, want = %q", got, want)
		}

	})

	t.Run("Not expected message", func(t *testing.T) {

		// Create a new instance of Person
		person := Person{
			FirstName: "Juan",
			LastName:  "Diaz",
		}

		// Marshal it into JSON prior to requesting
		personJSON, err := json.Marshal(person)

		// Make request with marshalled JSON as the POST body
		resp, err := http.Post(DefaultHTTPGetAddress, "application/json",
			bytes.NewBuffer(personJSON))

		if err != nil {
			t.Error("Could not make POST request to http")
		}

		// But for good measure, let's look at the response body.
		body, err := ioutil.ReadAll(resp.Body)

		var result ResponseBody
		err = json.Unmarshal([]byte(body), &result)

		if err != nil {
			t.Error("Error unmarshaling data from request.")
		}

		want := "Hello Juan Diaz"
		got := result.Message

		if got != want {

			res := fmt.Sprintf("got = %v, want = %v", got, want)

			t.Fatal(res)
			log.Fatal(res)
			t.Errorf("got = %q, want = %q", got, want)
		}

	})
}
