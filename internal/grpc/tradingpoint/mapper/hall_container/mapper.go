package hall_container

import (
	"fmt"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/hall_container"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
)

func ProtoTypeToModelType(t proto.HallContainerType) (model.Type, error) {
	const op = "grpc.tradingpoint.mapper.ProtoTypeToModelType"

	switch t {
	case proto.HallContainerType_HALL_CONTAINER_SECTION:
		return model.TypeSection, nil
	case proto.HallContainerType_HALL_CONTAINER_STORE:
		return model.TypeStore, nil
	default:
		return "", fmt.Errorf("%s: unknown type: %s", op, t)
	}
}

func ModelTypeToProtoType(t model.Type) (proto.HallContainerType, error) {
	const op = "grpc.tradingpoint.mapper.ModelTypeToProtoType"

	switch t {
	case model.TypeSection:
		return proto.HallContainerType_HALL_CONTAINER_SECTION, nil
	case model.TypeStore:
		return proto.HallContainerType_HALL_CONTAINER_STORE, nil
	default:
		return proto.HallContainerType_HALL_CONTAINER_STORE, fmt.Errorf("%s: unknown type: %s", op, t)
	}
}
