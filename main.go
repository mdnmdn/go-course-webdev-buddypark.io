package main

import (
	"fmt"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "<h1>Wow!</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
