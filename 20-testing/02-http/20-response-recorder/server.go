package main

import (
	"net/http"

	"github.com/exu/go-workshops/20-testing/02-http/20-response-recorder/handler"
)

func main() {
	http.HandleFunc("/", handler.Simple)
	http.ListenAndServe(":8080", nil)
}