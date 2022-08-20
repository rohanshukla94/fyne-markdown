package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppConfig struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var config AppConfig

func main() {

	//create fyne app

	a := app.New()

	//create window for app

	win := a.NewWindow("Markdown")

	//get ui
	edit, preview := config.makeUI()

	config.createMenuItems(win)
	//set content of window

	win.SetContent(container.NewHSplit(edit, preview))

	//show window and run app

	win.Resize(fyne.Size{Width: 1000, Height: 600})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *AppConfig) makeUI() (*widget.Entry, *widget.RichText) {

	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *AppConfig) createMenuItems(win fyne.Window) {

	openMenuItem := fyne.NewMenuItem("Open...", func() {

	})

	saveMenuItem := fyne.NewMenuItem("Save", func() {})

	saveAsMenuItem := fyne.NewMenuItem("Save as...", func() {})

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}
