package hall

import (
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/hall"
	hallcontainermapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/hall_container"
	tradingpointmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/trading_point"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
)

func ModelHallToProtoHall(h hall.Hall) (*proto.Hall, error) {
	const op = "grpc.tradingpoint.mapper.ModelHallToProtoHall"

	tradingPointType, err := tradingpointmapper.ModelTypeToProtoType(h.TradingPointType)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	hallContainerType, err := hallcontainermapper.ModelTypeToProtoType(h.HallContainerType)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	return &proto.Hall{
		Id:                h.ID,
		HallContainerId:   h.HallContainerID,
		HallContainerType: hallContainerType,
		TradingPointId:    h.TradingPointID,
		TradingPointType:  tradingPointType,
	}, nil
}
