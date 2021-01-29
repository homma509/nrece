package infra

import (
	"fmt"
	"log"

	"github.com/homma509/nrece/domain/model"
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
func (r *AppRepository) Get(id string) (app model.App, err error) {
	log.Println("[info] infra app get", id)
	if r.sql == nil {
		err = fmt.Errorf("SQL Handler is nil")
		return
	}
	err = r.sql.Where(&app, "id = ?", id)
	return
}
