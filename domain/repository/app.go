package repository

import (
	"github.com/homma509/nrece/domain/model"
)

// AppRepository ...
type AppRepository interface {
	Get(id string) (model.App, error)
}
