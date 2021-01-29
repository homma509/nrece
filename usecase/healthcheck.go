package usecase

import (
	"log"

	"github.com/homma509/nrece/domain/model"
)

// Healthcheck ...
type Healthcheck struct {
}

// NewHealthcheck ...
func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}

// Get ...
func (h *Healthcheck) Get() *model.Healthcheck {
	log.Println("[info] usecase healthcheck get")
	return &model.Healthcheck{
		Message: "OK",
	}
}
