package application

import (
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/repository/tradingpoint"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/application"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type application struct {
	ID               int32
	TradingPointID   int32
	TradingPointType string
	CreatedAt        time.Time
	OrderID          *int32
}

func (a application) ToProto() (*proto.Application, error) {
	const op = "repository.application.ToProto"

	protoTradingPointType, err := tradingpoint.StringToProtoTradingPointType(a.TradingPointType)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto trading point type: %w", op, err)
	}

	return &proto.Application{
		Id:               a.ID,
		TradingPointId:   a.TradingPointID,
		TradingPointType: protoTradingPointType,
		CreatedAt:        timestamppb.New(a.CreatedAt),
		OrderId:          a.OrderID,
	}, nil
}

func (a application) ToProtoWith(products []*proto.ProductApplication) (*proto.Application, error) {
	const op = "repository.application.ToProtoWith"

	protoApplication, err := a.ToProto()
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert application to proto: %w", op, err)
	}
	protoApplication.Products = products
	return protoApplication, nil
}

type productApplication struct {
	ID            int32
	ApplicationID int32
	ProductID     int32
	Quantity      int32
}

func (a productApplication) ToProto() *proto.ProductApplication {
	return &proto.ProductApplication{
		ProductId: a.ProductID,
		Quantity:  a.Quantity,
	}
}
