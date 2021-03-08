package value

// Phone ...
type Phone string

// NewPhone ...
func NewPhone(s string) (Phone, error) {
	return Phone(s), nil
}

func (p *Phone) String() string {
	return p.String()
}
