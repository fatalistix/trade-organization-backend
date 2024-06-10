package core

import (
	"fmt"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/core"
)

func ProtoDateToString(date *proto.Date) string {
	return fmt.Sprintf("%04d-%02d-%02d", date.Year, date.Month, date.Day)
}

func ProtoMoneyToString(money *proto.Money) string {
	return fmt.Sprintf("%d.%02d", money.Rubles, money.Pennies)
}

func StringToProtoMoney(money string) (*proto.Money, error) {
	const op = "repository.core.StringToProtoMoney"

	if len(money) == 0 {
		return &proto.Money{
			Rubles:  0,
			Pennies: 0,
		}, nil
	}

	var rubles int64
	var pennies int8
	_, err := fmt.Sscanf(money, "%d.%d", &rubles, &pennies)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to parse money: %w", op, err)
	}

	return &proto.Money{
		Rubles:  rubles,
		Pennies: int32(pennies),
	}, nil
}

func StringToProtoDate(date string) (*proto.Date, error) {
	const op = "repository.core.StringToProtoDate"

	var day, month, year int32
	_, err := fmt.Sscanf(date, "%d-%d-%d", &year, &month, &day)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to parse date: %w", op, err)
	}
	return &proto.Date{
		Year:  year,
		Month: month,
		Day:   day,
	}, nil
}
