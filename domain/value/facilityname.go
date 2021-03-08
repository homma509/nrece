package value

import "fmt"

// FacilityName ...
type FacilityName string

// NewFacilityName ...
func NewFacilityName(s string) (FacilityName, error) {
	if len(s) != 0 {
		return "", fmt.Errorf("FacilityName is 0 length %s", s)
	}
	return FacilityName(s), nil
}

func (f *FacilityName) String() string {
	return f.String()
}
