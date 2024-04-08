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
	w := a.NewWindow("Clothing Manager")

	w.Resize(fyne.NewSize(300, 400))
	data := readDataFile()

	clotingList := widget.NewList(
		// func that returns the number of items in the list
		func() int { return len(data) },

		// func that returns the component structure of the List Item
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil,
				container.NewBorder(
					nil, nil, widget.NewLabel(""), nil,
					widget.NewLabel(""),
				),
				nil,
				nil,
				container.NewBorder(
					nil, nil, widget.NewLabel(""),

					widget.NewCheck("", func(b bool) {}),
				),

				// takes the rest of the space
			)
		},
		// func that is called for each item in the list and allows
		// you to show the content on the previously defined ui structure
		func(i widget.ListItemID, o fyne.CanvasObject) {
			ctr, _ := o.(*fyne.Container)
			leftContainer := ctr.Objects[1].(*fyne.Container)
			rightContainer := ctr.Objects[0].(*fyne.Container)
			// ideally we should check `ok` for each one of those casting
			// but we know that they are those types for sure

			// For changing the layout, the take up rest space are counted as objects FIRST!
			typeLabel := leftContainer.Objects[0].(*widget.Label)
			colorLabel := leftContainer.Objects[1].(*widget.Label)

			descLabel := rightContainer.Objects[0].(*widget.Label)
			favCheck := rightContainer.Objects[1].(*widget.Check)

			typeLabel.SetText(data[i].Type)
			descLabel.SetText(data[i].Description)
			colorLabel.SetText(data[i].Color)
			favCheck.SetChecked(data[i].Fav)
		})

	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "New Clothing Item Description..."

	addBtn := widget.NewButton("Add", func() {
		data = addBtnFunc(data, newtodoDescTxt)
		newtodoDescTxt.Refresh()
		clotingList.Refresh()
	})
	addBtn.Disable()

	newtodoDescTxt.OnSubmitted = func(s string) {
		if addBtn.Disabled() {
			return
		}
		data = addBtnFunc(data, newtodoDescTxt)
		newtodoDescTxt.Refresh()
		clotingList.Refresh()
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
			clotingList,
		),
	)
	w.ShowAndRun()
}

func addBtnFunc(data []models.ClothingItem, entry *widget.Entry) []models.ClothingItem {
	addedTodo := models.NewItem("type", "color", entry.Text)
	updateDataFile(addedTodo)
	data = append(data, addedTodo)
	entry.Text = ""
	entry.OnChanged(entry.Text)
	return data
}

func updateDataFile(newData models.ClothingItem) {
	f, err := os.OpenFile("data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	bytes := newData.JSON()
	f.Write(append(bytes, 10))
}

func readDataFile() []models.ClothingItem {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var items []models.ClothingItem
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var item models.ClothingItem
		if err := json.Unmarshal(scanner.Bytes(), &item); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}
		items = append(items, item)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return items
}
