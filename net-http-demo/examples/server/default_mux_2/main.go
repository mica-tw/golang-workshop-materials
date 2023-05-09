package main

import (
	"fmt"
	"net/http"
	"sync"
)

type HitCounter struct {
	count int
	lock  sync.Mutex // Guard the counter
}

func (h *HitCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.count++
	fmt.Fprintf(w, "Number of hits: %d", h.count)
}

func main() {
	http.Handle("/", &HitCounter{})
	http.ListenAndServe(":8080", nil) // HL
}
