package main

import (
	"io"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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

	a.Settings().SetTheme(&myTheme{})

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

	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(win))

	saveMenuItem := fyne.NewMenuItem("Save", app.saveFunc(win))

	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (app *AppConfig) openFunc(win fyne.Window) func() {
	return func() {

		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {

			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				return
			}

			defer read.Close()

			data, err := io.ReadAll(read)

			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			app.EditWidget.SetText(string(data))

			app.CurrentFile = read.URI()

			win.SetTitle(win.Title() + " - " + read.URI().Name())

			app.SaveMenuItem.Disabled = false

		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *AppConfig) saveFunc(win fyne.Window) func() {
	return func() {

		if app.CurrentFile != nil {

			write, err := storage.Writer(app.CurrentFile)

			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			write.Write([]byte(app.EditWidget.Text))

			defer write.Close()
		}
	}
}

func (app *AppConfig) saveAsFunc(win fyne.Window) func() {

	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {

			if err != nil {
				dialog.ShowError(err, win)

				return
			}

			if write == nil {
				//user cancelled
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {

				dialog.ShowInformation("Error", "Please name your file with a .md ext", win)
				return

			}

			//save file
			write.Write([]byte(app.EditWidget.Text))

			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())

			app.SaveMenuItem.Disabled = false
		}, win)

		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}
