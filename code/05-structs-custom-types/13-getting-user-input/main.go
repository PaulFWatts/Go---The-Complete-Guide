package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(title, "\n", content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title: ")

	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note content: ")

	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	value = value[:len(value)-1] // Remove the trailing newline
	if value == "" {
		return "", errors.New("invalid input")
	}
	return value, nil
}
