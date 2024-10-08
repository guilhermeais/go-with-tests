package main

import (
	"go-with-tests/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now().Local()
	clockface.SVGWriter(os.Stdout, t)
}
