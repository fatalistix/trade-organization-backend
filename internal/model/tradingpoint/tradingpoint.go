package tradingpoint

import (
	"github.com/fatalistix/trade-organization-backend/internal/model/core"
	"time"
)

type TradingPoint struct {
	ID           int32
	Type         Type
	AreaPlot     float64
	RentalCharge *core.Money
	CounterCount int32
	Address      string
}

type UtilityService struct {
	ID         int32
	PaymentDay time.Time
	Amount     core.Money
}

type RentalCharge struct {
	ID         int32
	PaymentDay time.Time
	Amount     core.Money
}
