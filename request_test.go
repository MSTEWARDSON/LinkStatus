package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test200Response(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `response from the mock server goes here`)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - OK!"))
	}))
	defer ts.Close()

	mockServerURL := ts.URL

	checkStatus(mockServerURL, 1, false, 1)
}

func Test404Response(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `response from the mock server goes here`)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found!"))
	}))
	defer ts.Close()

	mockServerURL := ts.URL

	checkStatus(mockServerURL, 1, false, 1)
}

//300 status code check
func Test300Response(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `response from the mock server goes here`)
		w.WriteHeader(http.StatusMultipleChoices)
		w.Write([]byte("300 - Multiple choices!"))
	}))
	defer ts.Close()

	mockServerURL := ts.URL

	checkStatus(mockServerURL, 1, false, 1)
}
