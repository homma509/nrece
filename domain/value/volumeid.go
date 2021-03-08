package value

import (
	"fmt"
	"strconv"
)

// VolumeID ...
type VolumeID int

// NewVolumeID ...
func NewVolumeID(s string) (VolumeID, error) {
	if len(s) != 2 {
		return VolumeID(0), fmt.Errorf("VolumeID isn't 2 length %s", s)
	}
	no, err := strconv.Atoi(s)
	if err != nil {
		return VolumeID(0), fmt.Errorf("VolumeID couldn't convert numeric %s", s)
	}

	return VolumeID(no), nil
}

func (v *VolumeID) String() string {
	return fmt.Sprintf("%0d", v)
}
