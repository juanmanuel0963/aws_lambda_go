package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

var (
	DefaultHTTPAddress = "https://8syalbja7g.execute-api.us-east-1.amazonaws.com/hello_world"
)

func Test_hello_world(t *testing.T) {

	t.Run("GET_request", func(t *testing.T) {

		//GET request---------------------------------------------------

		resp, err := http.Get(DefaultHTTPAddress)

		if err != nil {
			message := "Unsuccessfull GET request to " + DefaultHTTPAddress
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
