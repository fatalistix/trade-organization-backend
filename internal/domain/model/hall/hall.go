package hall

import (
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/hall_container"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
)

type Hall struct {
	ID                int32
	HallContainerID   int32
	HallContainerType hall_container.Type
	TradingPointID    int32
	TradingPointType  trading_point.Type
}
