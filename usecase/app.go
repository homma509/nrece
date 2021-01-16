package usecase

import (
	"github.com/homma509/9go/domain/model"
	"github.com/homma509/9go/domain/repository"
)

// App ...
type App struct {
	repo repository.AppRepository
}

// NewApp ...
func NewApp(repo repository.AppRepository) *App {
	return &App{
		repo,
	}
}

// Get ...
func (a *App) Get(id string) (app model.App, err error) {
	app, err = a.repo.Get(id)
	return
}
