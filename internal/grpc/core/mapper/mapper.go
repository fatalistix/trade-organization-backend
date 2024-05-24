package core

import (
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/core"
)

func ModelMoneyToProtoMoney(money model.Money) *proto.Money {
	return &proto.Money{
		Rubles:  money.Rubles,
		Pennies: int32(money.Pennies),
	}
}

func ProtoMoneyToModelMoney(money *proto.Money) model.Money {
	return model.Money{
		Rubles:  money.Rubles,
		Pennies: int8(money.Pennies),
	}
}

func ModelDateToProtoDate(date model.Date) *proto.Date {
	return &proto.Date{
		Year:  date.Year,
		Month: date.Month,
		Day:   date.Day,
	}
}

func ProtoDateToModelDate(date *proto.Date) model.Date {
	return model.Date{
		Year:  date.Year,
		Month: date.Month,
		Day:   date.Day,
	}
}
