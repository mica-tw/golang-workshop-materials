package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest(
		"POST",
		"https://httpbingo.org/post",
		bytes.NewBuffer([]byte(`{"name": "Alice", "age": 25}`)),
	)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
