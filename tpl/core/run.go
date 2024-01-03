package core

import (
	"log"

	console "github.com/pastusiak/golang/tpl/console"
	coreconsole "github.com/pastusiak/golang/tpl/core_console"
	desktop "github.com/pastusiak/golang/tpl/desktop"
)

var ModeConsole string = "console"
var ModeDesktop string = "desktop"

var coreConsole coreconsole.App

type App struct {
	defaultMode       string
	enableGUIMode     bool
	enableConsoleMode bool
}

func NewApp(loadEnv bool, defaultMode string, enableGUIMode bool, enableConsoleMode bool) *App {
	coreConsole = *coreconsole.NewApp(loadEnv)

	a := new(App)
	a.defaultMode = defaultMode
	a.enableGUIMode = enableGUIMode
	a.enableConsoleMode = enableConsoleMode

	return a
}

func (a *App) RegisterCmd(cmd console.CmdInterface) {
	coreConsole.RegisterCmd(cmd)
}

func (a *App) RegisterWindow(win desktop.WindowInterface) {
	desktop.RegisterWindow(win)
}

func (a *App) Run() {
	defaultMode := a.defaultMode
	isManualMode, manualMode := console.ManualMode()

	if isManualMode {
		defaultMode = manualMode
	}

	if defaultMode == "desktop" {
		if a.enableGUIMode {
			desktop.Run()
		} else {
			log.Fatalf("GUI mode is disabled.")
		}
	} else {
		if a.enableConsoleMode {
			coreConsole.Run()
		} else {
			log.Fatalf("Console mode is disabled.")
		}
	}
}
