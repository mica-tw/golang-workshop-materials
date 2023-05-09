package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func greet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}
func main() {
	ts := httptest.NewServer(http.HandlerFunc(greet))
	defer ts.Close()
	res, _ := http.Get(ts.URL)
	if res.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}
