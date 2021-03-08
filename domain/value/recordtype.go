package value

import (
	"fmt"

	"github.com/homma509/nrece/pkg"
)

const (
	// IR ...
	IR string = "IR"
)

var recordTypes []string = []string{IR}

// RecordType ...
type RecordType string

// NewRecordType ...
func NewRecordType(s string) (RecordType, error) {
	if !pkg.Contains(s, recordTypes) {
		return "", fmt.Errorf("RecordType is invalid %s", s)
	}

	return RecordType(s), nil
}

func (r *RecordType) String() string {
	return r.String()
}
