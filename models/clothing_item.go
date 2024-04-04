package models

import (
	"encoding/json"
	"fmt"
)

type ClothingItem struct {
	Color       string
	Description string
	Fav         bool
}

func NewItem(color, description string) ClothingItem {
	return ClothingItem{color, description, false}
}
func (c ClothingItem) String() string {
	return fmt.Sprintf("%s - %s - %t", c.Color, c.Description, c.Fav)
}
func (c ClothingItem) JSON() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
