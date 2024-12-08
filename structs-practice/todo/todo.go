package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("not a valid input for a Todo")
	}

	return Todo{
		Content: content,
	}, nil
}

func (t Todo) Display() {
	fmt.Printf("Your Todo has the following content: %v\n", t.Content)
}

func (t Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0o644)
}
