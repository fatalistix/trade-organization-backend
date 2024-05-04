package tradingpoint

import (
	"fmt"
	modelcore "github.com/fatalistix/trade-organization-backend/internal/model/core"
	model "github.com/fatalistix/trade-organization-backend/internal/model/tradingpoint"
)

type tradingPoint struct {
	ID           int32
	Type         string
	AreaPlot     float64
	RentalCharge string
	CounterCount int32
	Address      string
}

func (t *tradingPoint) ToModel() (*model.TradingPoint, error) {
	const op = "repository.tradingpoint.ToModel"

	mt, err := model.TypeFromString(t.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to model type: %w", op, err)
	}
	mm, err := modelcore.MoneyFromString(t.RentalCharge)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to model money: %w", op, err)
	}
	return &model.TradingPoint{
		ID:           t.ID,
		Type:         mt,
		AreaPlot:     t.AreaPlot,
		RentalCharge: mm,
		CounterCount: t.CounterCount,
		Address:      t.Address,
	}, nil
}
