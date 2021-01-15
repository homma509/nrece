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
func (a *App) Get(id string) (*model.App, error) {
	res, err := a.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
