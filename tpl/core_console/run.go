package coreconsole

import (
	"github.com/joho/godotenv"
	"github.com/pastusiak/golang/tpl/console"
)

type App struct{}

func NewApp(loadEnv bool) *App {
	if loadEnv {
		godotenv.Load()
	}

	return &App{}
}

func (a *App) RegisterCmd(cmd console.CmdInterface) {
	console.RegisterCmd(cmd)
}

func (a *App) Run() {
	console.Run()
}
