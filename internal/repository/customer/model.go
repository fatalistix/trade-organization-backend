package customer

import (
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/repository/core"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/customer"
)

type customer struct {
	ID          int32
	FirstName   string
	LastName    string
	BirthDate   string
	PhoneNumber string
}

func (c customer) ToProto() (*proto.Customer, error) {
	birthDate, err := core.StringToProtoDate(c.BirthDate)
	if err != nil {
		return nil, fmt.Errorf("unable to convert string to proto date: %w", err)
	}

	return &proto.Customer{
		Id:          c.ID,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		BirthDate:   birthDate,
		PhoneNumber: c.PhoneNumber,
	}, nil
}
