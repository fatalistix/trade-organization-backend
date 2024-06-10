package mapper

import (
	"fmt"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/seller"
	grpccore "github.com/fatalistix/trade-organization-backend/internal/grpc/core/mapper"
	placeofworkmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/place_of_work"
	tradingpointmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/trading_point"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
)

func ModelSellerToProtoSeller(seller model.Seller) (*proto.Seller, error) {
	const op = "grpc.seller.mapper.modelSellerToProtoSeller"

	worksAt, err := ModelWorksAtToProtoWorksAt(seller.WorksAt)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	return &proto.Seller{
		Id:          seller.ID,
		FirstName:   seller.FirstName,
		LastName:    seller.LastName,
		MiddleName:  seller.MiddleName,
		BirthDate:   grpccore.ModelDateToProtoDate(seller.BirthDate),
		Salary:      grpccore.ModelMoneyToProtoMoney(seller.Salary),
		PhoneNumber: seller.PhoneNumber,
		WorksAt:     worksAt,
	}, nil
}

func ModelWorksAtToProtoWorksAt(worksAt *model.WorksAt) (*proto.WorksAt, error) {
	const op = "grpc.seller.mapper.modelPlaceOfWorkToProtoPlaceOfWork"

	if worksAt == nil {
		return nil, nil
	}

	tradingPointType, err := tradingpointmapper.ModelTypeToProtoType(worksAt.TradingPoint.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	placeOfWorkType, err := placeofworkmapper.ModelTypeToProtoType(worksAt.PlaceOfWork.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	return &proto.WorksAt{
		TradingPoint: &proto.TradingPoint{
			Id:   worksAt.TradingPoint.ID,
			Type: tradingPointType,
		},
		PlaceOfWork: &proto.PlaceOfWork{
			Id:   worksAt.PlaceOfWork.ID,
			Type: placeOfWorkType,
		},
	}, nil
}

func ProtoWorksAtToModelWorksAt(worksAt *proto.WorksAt) (*model.WorksAt, error) {
	const op = "grpc.seller.mapper.protoPlaceOfWorkToModelPlaceOfWork"

	if worksAt == nil {
		return nil, nil
	}

	tradingPointType, err := tradingpointmapper.ProtoTypeToModelType(worksAt.TradingPoint.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	placeOfWorkType, err := placeofworkmapper.ProtoTypeToModelType(worksAt.PlaceOfWork.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	return &model.WorksAt{
		TradingPoint: model.TradingPoint{
			ID:   worksAt.TradingPoint.Id,
			Type: tradingPointType,
		},
		PlaceOfWork: model.PlaceOfWork{
			ID:   worksAt.PlaceOfWork.Id,
			Type: placeOfWorkType,
		},
	}, nil
}
