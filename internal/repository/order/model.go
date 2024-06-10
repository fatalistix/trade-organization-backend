package order

import (
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/repository/core"
	"github.com/fatalistix/trade-organization-backend/internal/repository/tradingpoint"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type order struct {
	ID          int32
	SupplierID  int32
	CreatedAt   time.Time
	Status      string
	CompletedAt *time.Time
	CanceledAt  *time.Time
}

func (o order) ToProto() (*proto.Order, error) {
	const op = "repository.order.ToProto"

	status, err := StringToProtoOrderStatus(o.Status)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto order status: %w", op, err)
	}

	order := &proto.Order{
		Id:         o.ID,
		SupplierId: o.SupplierID,
		CreatedAt:  timestamppb.New(o.CreatedAt),
		Status:     status,
	}

	if o.CompletedAt != nil {
		order.CompletedAt = timestamppb.New(*o.CompletedAt)
	}

	if o.CanceledAt != nil {
		order.CanceledAt = timestamppb.New(*o.CanceledAt)
	}

	return order, nil
}

func (o order) ToProtoWith(products []*proto.ProductOrder, distributions []*proto.Distribution) (*proto.Order, error) {
	const op = "repository.order.ToProtoWith"

	protoOrder, err := o.ToProto()
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert order to proto: %w", op, err)
	}

	protoOrder.Products = products
	protoOrder.Distributions = distributions

	return protoOrder, nil
}

type productOrder struct {
	ID        int32
	OrderID   int32
	ProductID int32
	Quantity  int32
	Price     string
}

func (o productOrder) ToProto() (*proto.ProductOrder, error) {
	const op = "repository.order.ToProto"

	price, err := core.StringToProtoMoney(o.Price)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto money: %w", op, err)
	}

	return &proto.ProductOrder{
		ProductId: o.ProductID,
		Quantity:  o.Quantity,
		Price:     price,
	}, nil
}

type distribution struct {
	ID               int32
	TradingPointID   int32
	TradingPointType string
	OrderID          int32
	ProductID        int32
	Quantity         int32
}

func (d distribution) ToProto() (*proto.Distribution, error) {
	const op = "repository.order.ToProto"

	protoTradingPointType, err := tradingpoint.StringToProtoTradingPointType(d.TradingPointType)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto trading point type: %w", op, err)
	}

	return &proto.Distribution{
		TradingPointId:   d.TradingPointID,
		TradingPointType: protoTradingPointType,
		ProductId:        d.ProductID,
		Quantity:         d.Quantity,
	}, nil
}
