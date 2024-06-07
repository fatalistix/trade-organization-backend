package supplier

import (
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/repository/core"
	protoproduct "github.com/fatalistix/trade-organization-proto/gen/go/product"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/supplier"
)

type supplier struct {
	ID   int32
	Name string
	Type string
}

func (s supplier) ToProtoWith(products []*proto.ProductSupplier) (*proto.Supplier, error) {
	const op = "repository.supplier.ToProto"

	protoType, err := StringToProtoSupplierType(s.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto type: %w", op, err)
	}
	return &proto.Supplier{
		Id:       s.ID,
		Name:     s.Name,
		Type:     protoType,
		Products: products,
	}, nil
}

func (s supplier) ToProto() (*proto.Supplier, error) {
	const op = "repository.supplier.ToModel"

	protoType, err := StringToProtoSupplierType(s.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto type: %w", op, err)
	}
	return &proto.Supplier{
		Id:   s.ID,
		Name: s.Name,
		Type: protoType,
	}, nil
}

type productSupplier struct {
	ID         int32
	SupplierID int32
	ProductID  int32
	Price      string
}

func (s productSupplier) ToProtoWith(products []*protoproduct.Product) (*proto.ProductSupplier, error, bool) {
	const op = "repository.supplier.productSupplier.ToProto"

	price, err := core.StringToProtoMoney(s.Price)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto money: %w", op, err), false
	}

	protoProductSupplier := &proto.ProductSupplier{
		Price: price,
	}

	for i := range products {
		if products[i].Id == s.ProductID {
			protoProductSupplier.Product = products[i]
			return protoProductSupplier, nil, true
		}
	}

	return nil, fmt.Errorf("%s: product not found", op), false
}
