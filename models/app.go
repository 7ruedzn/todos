package models

type App struct {
	Config Config
	Todos  []Todo
}

type Config struct {
	Path string
}

func (a *App) GetApp() *App {
	return a
}

func New() *App {
	return &App{}
}
