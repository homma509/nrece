package repository

import (
	"github.com/homma509/9go/domain/model"
)

// AppRepository ...
type AppRepository interface {
	Get(id string) (model.App, error)
}
