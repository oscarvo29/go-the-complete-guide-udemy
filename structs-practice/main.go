package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ovo/structs-practice/note"
	"ovo/structs-practice/todo"
)

type Saver interface {
	Save() error
}

type outputtable interface {
	Saver
	Display()
}

func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	err := saveData(data)
	if err != nil {
		return err
	}

	return nil
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the succeeded!")
	return err
}

func getUserInput(prompt string) string {
	fmt.Print(prompt + "\t")

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}
