package value

import (
	"fmt"
)

var pointTables map[string]string = map[string]string{
	"1": "医科",
}

// PointTable ...
type PointTable struct {
	ID   string
	Name string
}

// NewPointTable ...
func NewPointTable(id string) (PointTable, error) {
	name, ok := pointTables[id]
	if !ok {
		return PointTable{}, fmt.Errorf("PointTable is invalid %s", id)
	}
	return PointTable{
		ID:   id,
		Name: name,
	}, nil
}

func (p *PointTable) String() string {
	return p.Name
}
