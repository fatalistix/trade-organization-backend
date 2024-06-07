package product

import proto "github.com/fatalistix/trade-organization-proto/gen/go/product"

type product struct {
	ID          int32
	Name        string
	Description string
}

func (p product) ToProto() *proto.Product {
	return &proto.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}
}
