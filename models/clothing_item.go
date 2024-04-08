package models

import (
	"encoding/json"
	"fmt"
)

type ClothingItem struct {
	Type        string //maybe i can change this into a struct or interface later?
	Color       string
	Description string
	Fav         bool
}

func NewItem(typ, color, description string) ClothingItem {
	return ClothingItem{typ, color, description, false}
}
func (c ClothingItem) String() string {
	return fmt.Sprintf("%s - %s - %t", c.Color, c.Description, c.Fav)
}
func (t ClothingItem) JSON() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
