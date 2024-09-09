package main

import (
	"go-with-tests/dependencyinjection"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreetHandler)))
}
