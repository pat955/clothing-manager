package main

import (
	"clothing_manager/models"
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
	data := []models.Todo{
		models.NewTodo("Some stuff"),
	}

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
		addedTodo := models.NewTodo(newtodoDescTxt.Text)
		updateDataFile(addedTodo)
		data = append(data, addedTodo)
		todoList.Refresh()
		newtodoDescTxt.Text = ""
		newtodoDescTxt.Refresh()
		newtodoDescTxt.OnChanged(newtodoDescTxt.Text)
	})
	addBtn.Disable()

	newtodoDescTxt.OnSubmitted = func(s string) {
		if addBtn.Disabled() {
			return
		}
		addedTodo := models.NewTodo(newtodoDescTxt.Text)
		updateDataFile(addedTodo)
		data = append(data, addedTodo)
		todoList.Refresh()
		newtodoDescTxt.Text = ""
		newtodoDescTxt.Refresh()
		newtodoDescTxt.OnChanged(newtodoDescTxt.Text)
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

func updateDataFile(newData models.Todo) {
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(newData.String() + "\n")
}
