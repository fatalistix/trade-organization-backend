package core

import "fmt"

type Date struct {
	Year  int32
	Month int32
	Day   int32
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Day, d.Month)
}

func DateFromString(date string) (Date, error) {
	const op = "domain.model.core.DateFromString"

	d := Date{}

	_, err := fmt.Sscanf(date, "%04d-%02d-%02d", &d.Year, &d.Day, &d.Month)
	if err != nil {
		return Date{}, fmt.Errorf("%s: %w", op, err)
	}

	return d, nil
}
