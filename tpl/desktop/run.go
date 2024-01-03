package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var winList []WindowInterface
var WindowApp fyne.App

func RegisterWindow(win WindowInterface) {
	winList = append(winList, win)
}

func Run() bool {
	WindowApp = app.New()
	fyne.SetCurrentApp(WindowApp)

	for _, item := range winList {
		window := item.Window()
		window.Show()
	}

	WindowApp.Run()

	return true
}
