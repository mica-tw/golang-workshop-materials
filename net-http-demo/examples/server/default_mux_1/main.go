package main

import "net/http"

type webHandler string

func (wh webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(wh))
}

func main() {
	http.Handle("/api/", webHandler("API Response"))
	http.Handle("/web/", webHandler("WEB Response"))
	http.ListenAndServe(":8080", nil) // HL
}
