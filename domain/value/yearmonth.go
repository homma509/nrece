package value

import (
	"fmt"
	"strconv"
)

// YearMonth ...
type YearMonth struct {
	Year  int
	Month int
}

// NewYearMonth ...
func NewYearMonth(s string) (YearMonth, error) {
	if len(s) != 6 {
		return YearMonth{}, fmt.Errorf("YearMonth isn't 6 length %s", s)
	}
	y, err := strconv.Atoi(s[:4])
	if err != nil {
		return YearMonth{}, fmt.Errorf("YearMonth couldn't convert year %s", s)
	}
	m, err := strconv.Atoi(s[4:])
	if err != nil {
		return YearMonth{}, fmt.Errorf("YearMonth couldn't convert month %s", s)
	}
	if m < 0 || 12 < m {
		return YearMonth{}, fmt.Errorf("YearMonth couldn't convert month %s", s)
	}

	return YearMonth{
		Year:  y,
		Month: m,
	}, nil
}

func (ym *YearMonth) String() string {
	return fmt.Sprintf("%d-%d", ym.Year, ym.Month)
}
