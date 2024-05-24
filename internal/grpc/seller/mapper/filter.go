package mapper

import (
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/place_of_work"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/seller"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
	placeofworkmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/place_of_work"
	tradingpointmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/trading_point"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
)

func ListRequestToModelFilter(request *proto.ListRequest) (*model.Filter, error) {
	const op = "grpc.seller.mapper.ListRequestToModelFilter"

	worksAtFilterType, err := ProtoWorksAtFilterTypeToModelWorksAtFilterType(request.WorksAtFilter)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	var tradingPointType *trading_point.Type
	if request.TradingPointType != nil {
		tradingPointTypeTemp, err := tradingpointmapper.ProtoTypeToModelType(*request.TradingPointType)
		if err != nil {
			return nil, fmt.Errorf("%s: error mapping: %w", op, err)
		}
		tradingPointType = &tradingPointTypeTemp
	}

	var placeOfWorkType *place_of_work.Type
	if request.PlaceOfWorkType != nil {
		placeOfWorkTypeTemp, err := placeofworkmapper.ProtoTypeToModelType(*request.PlaceOfWorkType)
		if err != nil {
			return nil, fmt.Errorf("%s: error mapping: %w", op, err)
		}
		placeOfWorkType = &placeOfWorkTypeTemp
	}

	return &model.Filter{
		WorksAtFilterType: worksAtFilterType,
		TradingPointId:    request.TradingPointId,
		TradingPointType:  tradingPointType,
		PlaceOfWorkId:     request.PlaceOfWorkId,
		PlaceOfWorkType:   placeOfWorkType,
		Search:            request.Search,
	}, nil
}

func ProtoWorksAtFilterTypeToModelWorksAtFilterType(t proto.WorksAtFilterType) (model.WorksAtFilterType, error) {
	const op = "grpc.seller.mapper.ProtoWorksAtTypeToModelWorksAtType"

	switch t {
	case proto.WorksAtFilterType_WORKS_AT_FILTER_TYPE_ALL:
		return model.WorksAtFilterTypeAll, nil
	case proto.WorksAtFilterType_WORKS_AT_FILTER_TYPE_ONLY_NULL:
		return model.WorksAtFilterTypeOnlyNull, nil
	case proto.WorksAtFilterType_WORKS_AT_FILTER_TYPE_ONLY_NOT_NULL:
		return model.WorksAtFilterTypeOnlyNotNull, nil
	default:
		return "", fmt.Errorf("%s: unknown type: %s", op, t)
	}
}
