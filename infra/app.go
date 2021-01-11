package infra

import "github.com/homma509/9go/domain/model"

// AppRepository ...
type AppRepository struct {
	*SQLHandler
}

// NewAppRepository ...
func NewAppRepository(sqlHandler *SQLHandler) *AppRepository {
	return &AppRepository{
		sqlHandler,
	}
}

// Get ...
func (r *AppRepository) Get(id string) (app *model.App, err error) {
	err = r.Where(&app, "id = ?", id)
	return
}
