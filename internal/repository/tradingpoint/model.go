package tradingpoint

import (
	"fmt"
	modelcore "github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	modelhall "github.com/fatalistix/trade-organization-backend/internal/domain/model/hall"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/hall_container"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/receipting_point_with_accounting"
	modelsection "github.com/fatalistix/trade-organization-backend/internal/domain/model/section"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
)

type tradingPoint struct {
	ID           int32
	Type         string
	AreaPlot     float64
	RentalCharge string
	CounterCount int32
	Address      string
}

func (t tradingPoint) ToModel() (trading_point.TradingPoint, error) {
	const op = "repository.tradingpoint.ToModel"

	mt, err := trading_point.TypeFromString(t.Type)
	if err != nil {
		return trading_point.TradingPoint{}, fmt.Errorf("%s: unable to convert string to model type: %w", op, err)
	}
	mm, err := modelcore.MoneyFromString(t.RentalCharge)
	if err != nil {
		return trading_point.TradingPoint{}, fmt.Errorf("%s: unable to convert string to model money: %w", op, err)
	}
	return trading_point.TradingPoint{
		ID:           t.ID,
		Type:         mt,
		AreaPlot:     t.AreaPlot,
		RentalCharge: mm,
		CounterCount: t.CounterCount,
		Address:      t.Address,
	}, nil
}

type receiptingPointWithAccounting struct {
	ID   int32
	Type string
}

func (r receiptingPointWithAccounting) ToModel() (receipting_point_with_accounting.ReceiptingPointWithAccounting, error) {
	const op = "repository.tradingpoint.ToModel"

	mt, err := receipting_point_with_accounting.TypeFromString(r.Type)
	if err != nil {
		return receipting_point_with_accounting.ReceiptingPointWithAccounting{}, fmt.Errorf("%s: unable to convert string to model type: %w", op, err)
	}

	return receipting_point_with_accounting.ReceiptingPointWithAccounting{
		ID:   r.ID,
		Type: mt,
	}, nil
}

type hall struct {
	ID                int32
	Type              string
	HallContainerID   int32
	HallContainerType string
	TradingPointID    int32
	TradingPointType  string
}

func (h hall) ToModel() (modelhall.Hall, error) {
	const op = "repository.tradingpoint.ToModel"

	hct, err := hall_container.TypeFromString(h.HallContainerType)
	if err != nil {
		return modelhall.Hall{}, fmt.Errorf("%s: unable to convert string to model type: %w", op, err)
	}

	tpt, err := trading_point.TypeFromString(h.TradingPointType)
	if err != nil {
		return modelhall.Hall{}, fmt.Errorf("%s: unable to convert string to model type: %w", op, err)
	}

	return modelhall.Hall{
		ID:                h.ID,
		HallContainerID:   h.HallContainerID,
		HallContainerType: hct,
		TradingPointID:    h.TradingPointID,
		TradingPointType:  tpt,
	}, nil
}

type section struct {
	ID                int32
	Type              string
	DepartmentStoreID int32
}

func (s section) ToModel() modelsection.Section {
	return modelsection.Section{
		ID:                s.ID,
		DepartmentStoreID: s.DepartmentStoreID,
	}
}
