package main

import "fmt"

func Hello(_name ...string) string {
	name := "World"

	if len(_name) > 0 {
		for i, n := range _name {
			if i == 0 {
				name = n
			} else {
				name += " " + n
			}
		}
	}

	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello())
}
