package main

import (

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "clothing_manager/models"
)

func main() {
    a := app.New()
    w := a.NewWindow("TODO App")

        // ADDING THIS HERE
    w.Resize(fyne.NewSize(300, 400))
    t := models.NewTodo("Show this on the window")

    w.SetContent(
        container.NewCenter(
            widget.NewLabel(t.String()),
        ),
    )
    w.ShowAndRun()
}