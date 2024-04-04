package main

import (
	"bufio"
	"clothing_manager/models"
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")

	w.Resize(fyne.NewSize(300, 400))
	data := readDataFile()

	todoList := widget.NewList(
		// func that returns the number of items in the list
		func() int {
			return len(data)
		},
		// func that returns the component structure of the List Item
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil, nil, nil,
				// left of the border
				widget.NewCheck("", func(b bool) {}),
				// takes the rest of the space
				widget.NewLabel(""),
			)
		},
		// func that is called for each item in the list and allows
		// you to show the content on the previously defined ui structure
		func(i widget.ListItemID, o fyne.CanvasObject) {
			ctr, _ := o.(*fyne.Container)
			// ideally we should check `ok` for each one of those casting
			// but we know that they are those types for sure
			l := ctr.Objects[0].(*widget.Label)
			c := ctr.Objects[1].(*widget.Check)
			l.SetText(data[i].Description)
			c.SetChecked(data[i].Done)
		})

	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "New Todo Description..."

	addBtn := widget.NewButton("Add", func() {
		data = addBtnFunc(data, newtodoDescTxt)
		newtodoDescTxt.Refresh()
		todoList.Refresh()
	})
	addBtn.Disable()

	newtodoDescTxt.OnSubmitted = func(s string) {
		if addBtn.Disabled() {
			return
		}
		data = addBtnFunc(data, newtodoDescTxt)
		newtodoDescTxt.Refresh()
		todoList.Refresh()
	}

	newtodoDescTxt.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	addButtonArea := container.NewBorder(
		nil,            // TOP
		nil,            // BOTTOM
		nil,            // LEFT
		addBtn,         // RIGHT
		newtodoDescTxt, // REST
	)

	w.SetContent(
		container.NewBorder(
			nil,           // TOP
			addButtonArea, // BOTTOM
			nil,           // RIGHT
			nil,           // LEFT
			// the rest will take all the rest of the space
			todoList,
		),
	)
	w.ShowAndRun()
}

func addBtnFunc(data []models.Todo, entry *widget.Entry) []models.Todo {
	addedTodo := models.NewTodo(entry.Text)
	updateDataFile(addedTodo)
	data = append(data, addedTodo)
	entry.Text = ""
	entry.OnChanged(entry.Text)
	return data
}

func updateDataFile(newData models.Todo) {
	f, err := os.OpenFile("data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	bytes := newData.JSON()
	f.Write(append(bytes, 10))
}

func readDataFile() []models.Todo {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var todos []models.Todo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var todo models.Todo
		if err := json.Unmarshal(scanner.Bytes(), &todo); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}
		todos = append(todos, todo)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return todos
}
