package place_of_work

import (
	"fmt"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/place_of_work"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
)

func ModelTypeToProtoType(mt model.Type) (proto.PlaceOfWorkType, error) {
	const op = "grpc.tradingpoint.mapper.placeofwork.ModelTypeToProtoType"

	switch mt {
	case model.TypeHall:
		return proto.PlaceOfWorkType_PLACE_OF_WORK_HALL, nil
	case model.TypeKiosk:
		return proto.PlaceOfWorkType_PLACE_OF_WORK_KIOSK, nil
	case model.TypeTray:
		return proto.PlaceOfWorkType_PLACE_OF_WORK_KIOSK, nil
	default:
		return proto.PlaceOfWorkType_PLACE_OF_WORK_HALL, fmt.Errorf("%s: unknown type: %s", op, mt)
	}
}

func ProtoTypeToModelType(pt proto.PlaceOfWorkType) (model.Type, error) {
	const op = "grpc.tradingpoint.mapper.placeofwork.ProtoTypeToModelType"

	switch pt {
	case proto.PlaceOfWorkType_PLACE_OF_WORK_HALL:
		return model.TypeHall, nil
	case proto.PlaceOfWorkType_PLACE_OF_WORK_KIOSK:
		return model.TypeKiosk, nil
	case proto.PlaceOfWorkType_PLACE_OF_WORK_TRAY:
		return model.TypeTray, nil
	default:
		return "", fmt.Errorf("%s: unknown type: %s", op, pt)
	}
}
