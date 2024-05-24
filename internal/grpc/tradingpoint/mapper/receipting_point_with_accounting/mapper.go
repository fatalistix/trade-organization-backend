package receipting_point_with_accounting

import (
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/receipting_point_with_accounting"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
)

func ModelReceiptingPointWithAccountingToProtoReceiptingPointWithAccounting(r receipting_point_with_accounting.ReceiptingPointWithAccounting) (*proto.ReceiptingPointWithAccounting, error) {
	const op = "grpc.tradingpoint.mapper.ModelReceiptingPointWithAccountingToProtoReceiptingPointWithAccounting"

	receiptingPointWithAccountingType, err := ModelReceiptingPointWithAccountingTypeToProtoReceiptingPointWithAccountingType(r.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}

	return &proto.ReceiptingPointWithAccounting{
		Id:   r.ID,
		Type: receiptingPointWithAccountingType,
	}, nil
}

func ModelReceiptingPointWithAccountingTypeToProtoReceiptingPointWithAccountingType(t receipting_point_with_accounting.Type) (proto.ReceiptingPointWithAccountingType, error) {
	const op = "grpc.tradingpoint.mapper.ModelReceiptingPointWithAccountingTypeToProtoReceiptingPointWithAccountingType"

	switch t {
	case receipting_point_with_accounting.TypeDepartmentStore:
		return proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_DEPARTMENT_STORE, nil
	case receipting_point_with_accounting.TypeStore:
		return proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_STORE, nil
	default:
		return proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_STORE, fmt.Errorf("%s: unknown type: %s", op, t)
	}
}
