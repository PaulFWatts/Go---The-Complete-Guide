package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

// slowGreet simulates a long-running task by sleeping for 3 seconds
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!")
	fmt.Println(<-done)
}
