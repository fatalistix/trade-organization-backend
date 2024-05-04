package tradingpoint

import (
	"fmt"
	model "github.com/fatalistix/trade-organization-backend/internal/model/tradingpoint"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
)

func ProtoTypeToModelType(t proto.TradingPointType) (model.Type, error) {
	const op = "grpc.tradingpoint.protoTypeToModelType"

	switch t {
	case proto.TradingPointType_TRADING_POINT_DEPARTMENT_STORE:
		return model.TypeDepartmentStore, nil
	case proto.TradingPointType_TRADING_POINT_STORE:
		return model.TypeStore, nil
	case proto.TradingPointType_TRADING_POINT_KIOSK:
		return model.TypeKiosk, nil
	case proto.TradingPointType_TRADING_POINT_TRAY:
		return model.TypeTray, nil
	default:
		return "", fmt.Errorf("%s: unknown type: %s", op, t)
	}
}

func ModelTypeToProtoType(t model.Type) (proto.TradingPointType, error) {
	const op = "grpc.tradingpoint.modelTypeToProtoType"

	switch t {
	case model.TypeDepartmentStore:
		return proto.TradingPointType_TRADING_POINT_DEPARTMENT_STORE, nil
	case model.TypeStore:
		return proto.TradingPointType_TRADING_POINT_STORE, nil
	case model.TypeKiosk:
		return proto.TradingPointType_TRADING_POINT_KIOSK, nil
	case model.TypeTray:
		return proto.TradingPointType_TRADING_POINT_TRAY, nil
	default:
		return proto.TradingPointType_TRADING_POINT_STORE, fmt.Errorf("%s: unknown type: %s", op, t)
	}
}
