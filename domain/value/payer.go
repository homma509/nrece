package value

import (
	"fmt"
)

var payers map[string]string = map[string]string{
	"1": "社会保険診療報酬支払基金",
	"2": "国民健康保険団体連合会",
}

// Payer ...
type Payer struct {
	ID   string
	Name string
}

// NewPayer ...
func NewPayer(id string) (Payer, error) {
	name, ok := payers[id]
	if !ok {
		return Payer{}, fmt.Errorf("Payer is invalid %s", id)
	}

	return Payer{
		ID:   id,
		Name: name,
	}, nil
}

func (p *Payer) String() string {
	return p.Name
}
