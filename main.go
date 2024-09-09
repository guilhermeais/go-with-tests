package main

import (
	"go-with-tests/counter"
	"go-with-tests/dependencyinjection"
	"log"
	"net/http"
	"os"
)

func dependencyInjectionExample() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreetHandler)))
}

func countdownExample() {
	counter.Countdown(os.Stdout)
}
func main() {
	go dependencyInjectionExample()
	countdownExample()
}
