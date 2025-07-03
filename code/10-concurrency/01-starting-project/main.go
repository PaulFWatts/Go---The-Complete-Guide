// This is a simple program that demonstrates how to greet users
package main

import (
	"fmt"
	"time"
)

// This function greets the user with a given phrase
func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

// This function simulates a slow greeting by sleeping for 3 seconds
func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
}

func main() {
	greet("Nice to meet you!")
	greet("How are you?")
	slowGreet("How ... are ... you ...?")
	greet("I hope you're liking the course!")
}
