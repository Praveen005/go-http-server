package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T){
	tests := []struct{
		name		string
		path		string
		expected	string
	}{
		{
			name:		"index",
			path:		"/api",
			expected: 	"Hello World!",
		},
		{
			name:		"healthcheck",
			path:		"/healthz",
			expected: 	"ok",
		},
	}
	// http.NewServeMux() creates a new instance of a ServeMux type. ServeMux is short for "Serve Multiplexer," and it is an HTTP request multiplexer.
	// The ServeMux type is used to route incoming HTTP requests to their respective handlers based on the request's URL path.
	mux := http.NewServeMux()
	setupHandlers(mux)

	ts := httptest.NewServer(mux)
	defer ts.Close()
	for _, tc := range tests{
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(ts.URL + tc.path)
			if err != nil{
				t.Fatal(err)
			}
			respBody, err := io.ReadAll(resp.Body)
			if err != nil{
				t.Fatal(err)
			}
			if string(respBody) != tc.expected{
				t.Errorf(
					"Expected: %s, Got: %s",
					tc.expected, string(respBody),
				)
			}
		})
	}
}