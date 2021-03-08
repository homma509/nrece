package value

// Reserve ...
type Reserve string

// NewReserve ...
func NewReserve(s string) (Reserve, error) {
	return Reserve(s), nil
}

func (r *Reserve) String() string {
	return r.String()
}
