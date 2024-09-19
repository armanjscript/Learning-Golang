package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// with struct tag (metadata for each field)
type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	//write to file as json
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
