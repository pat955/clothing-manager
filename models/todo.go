package models

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	Description string
	Done        bool
}

func NewTodo(description string) Todo {
	return Todo{description, false}
}
func (t Todo) String() string {
	return fmt.Sprintf("%s - %t", t.Description, t.Done)
}
func (t Todo) JSON() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
