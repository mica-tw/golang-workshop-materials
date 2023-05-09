package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const data = `{ "lat": 13.0827, "long": 80.2707}`

func main() {
	resp, _ := http.Post("https://httpbingo.org/post", "application/json", strings.NewReader(data)) // HL
	//handle error
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// handle error
	fmt.Println("Status:", resp.Status, "Body:", string(body))

}
