package value

import "fmt"

// FacilityID ...
type FacilityID string

// NewFacilityID ...
func NewFacilityID(s string) (FacilityID, error) {
	if len(s) != 7 {
		return "", fmt.Errorf("FacilityID isn't 7 length %s", s)
	}
	return FacilityID(s), nil
}

func (f *FacilityID) String() string {
	return f.String()
}
