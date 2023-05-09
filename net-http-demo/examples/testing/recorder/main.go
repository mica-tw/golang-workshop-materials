package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func handlerToTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Found"))
}
func main() {
	req, _ := http.NewRequest("GET", "/", nil)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerToTest)
	handler.ServeHTTP(recorder, req)

	expected := "OK"
	if recorder.Body.String() != expected {
		fmt.Printf("ERROR: unexpected body: got %v want %v", recorder.Body.String(), expected)
	}
}
