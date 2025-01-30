package models

type App struct {
	Config Config
	Todos  []Todo
}

type Config struct {
	Path string
}
