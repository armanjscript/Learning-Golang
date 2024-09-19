package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// creating interface for saving data to file - (note, todo)
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// combine interfaces together
// type outputable interface {
// 	Save() error
// 	Display()
// }

// embedded interface - better solution for combining
type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	printSomething(todo)
}

func printSomething(value interface{}) {

	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	}

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// function for implementing interface
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
