package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"sync"
)

const (
	winTitle  = "Download Manager"
	winWidth  = 1240
	winHeight = 720
)

var terminated int
var wg sync.WaitGroup

type download struct {
	url    string
	path   string
	size   int
	status int
}

func main() {

	terminated = 0
	//tcount := 3

	app := app.New()

	w := app.NewWindow(winTitle)
	w.Resize(fyne.NewSize(winWidth, winHeight))

	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		log.Println("tapped home")
	})

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := container.New(layout.NewBorderLayout(bar, nil, nil, nil), bar, addButton)

	w.SetContent(content)

	log.Println("Starting GUI...")
	w.ShowAndRun()

	//q := NewQ(1000)
	//for i := 0; i < tcount; i++ {
	//	wg.Add(1)
	//}
}
