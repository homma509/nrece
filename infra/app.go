package infra

import (
	"fmt"

	"github.com/homma509/9go/domain/model"
)

// AppRepository ...
type AppRepository struct {
	sql *SQLHandler
}

// NewAppRepository ...
func NewAppRepository(sqlHandler *SQLHandler) *AppRepository {
	return &AppRepository{
		sqlHandler,
	}
}

// Get ...
func (r *AppRepository) Get(id string) (app *model.App, err error) {
	if r.sql == nil {
		err = fmt.Errorf("SQL Handler is nil")
		return
	}
	err = r.sql.Where(&app, "id = ?", id)
	return
}
