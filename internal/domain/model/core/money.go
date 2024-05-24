package core

import "fmt"

type Money struct {
	Rubles  int64
	Pennies int8
}

func (m Money) String() string {
	return fmt.Sprintf("%d.%d", m.Rubles, m.Pennies)
}

func MoneyFromString(str string) (Money, error) {
	const op = "model.core.MoneyFromString"

	var rubles int64
	var pennies int8
	_, err := fmt.Sscanf(str, "%d.%d", &rubles, &pennies)
	if err != nil {
		return Money{}, fmt.Errorf("%s: unable to parse money: %w", op, err)
	}

	return Money{
		Rubles:  rubles,
		Pennies: pennies,
	}, nil
}
