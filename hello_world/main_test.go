package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

var (
	DefaultHTTPGetAddress = "https://8syalbja7g.execute-api.us-east-1.amazonaws.com/hello_world"
)

func Test_hello_world(t *testing.T) {

	t.Run("Invalid URL", func(t *testing.T) {

		_, err := http.Get(DefaultHTTPGetAddress)

		if err != nil {
			t.Fatal("Invalid URL")
			log.Fatal("Invalid URL")
			t.Errorf("Invalid URL")

		}
	})

	t.Run("Non 200 response", func(t *testing.T) {

		want := 200
		resp, _ := http.Get(DefaultHTTPGetAddress)
		got := resp.StatusCode

		if got != want {

			res := fmt.Sprintf("got = %v, want = %v", got, want)

			t.Fatal(res)
			log.Fatal(res)
			t.Errorf("got = %q, want = %q", got, want)
		}

	})
}
