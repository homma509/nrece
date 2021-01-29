package usecase

import (
	"log"

	"github.com/homma509/nrece/domain/model"
	"github.com/homma509/nrece/domain/repository"
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
	log.Println("[info] usecase app get", id)
	app, err = a.repo.Get(id)
	return
}
