package core

import (
	model "github.com/fatalistix/trade-organization-backend/internal/model/core"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/core"
)

func ModelMoneyToProtoMoney(money *model.Money) *proto.Money {
	return &proto.Money{
		Rubles:  money.Rubles,
		Pennies: int32(money.Pennies),
	}
}

func ProtoMoneyToProtoMoney(money *proto.Money) *model.Money {
	return &model.Money{
		Rubles:  money.Rubles,
		Pennies: int8(money.Pennies),
	}
}
