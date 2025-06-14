package main

import (
	"bufio"
	"fmt"
	"os"

	"example14.com/note/note"
)

func main() {
	title, content := getNoteData()

	_, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	if len(value) > 0 && value[len(value)-1] == '\n' {
		value = value[:len(value)-1]
	}
	return value
}
