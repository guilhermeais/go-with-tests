package main

import (
	"go-with-tests/counter"
	"go-with-tests/dependencyinjection"
	"log"
	"net/http"
	"os"
	"time"
)

func dependencyInjectionExample() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreetHandler)))
}

func countdownExample() {
	sleeper := counter.MakeConfigurableSleeper(1*time.Second, time.Sleep)
	counter.Countdown(os.Stdout, sleeper)
}
func main() {
	go dependencyInjectionExample()
	countdownExample()
}
