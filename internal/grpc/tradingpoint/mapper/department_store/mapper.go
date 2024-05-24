package department_store

import (
	"fmt"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/department_store"
	hallmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/hall"
	receiptingpointwithaccountingmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/receipting_point_with_accounting"
	sectionmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/section"
	tradingpointmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/trading_point"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
)

func ModelDepartmentStoreToProtoDepartmentStore(ds model.DepartmentStore) (*proto.DepartmentStore, error) {
	tradingPoint, err := tradingpointmapper.ModelTradingPointToProtoTradingPoint(ds.TradingPoint)
	if err != nil {
		return nil, fmt.Errorf("unable to map trading point: %w", err)
	}

	receiptingPointWithAccounting, err := receiptingpointwithaccountingmapper.ModelReceiptingPointWithAccountingToProtoReceiptingPointWithAccounting(ds.ReceiptingPointWithAccounting)
	if err != nil {
		return nil, fmt.Errorf("unable to map receipting point with accounting: %w", err)
	}

	sections := make([]*proto.Section, 0, len(ds.Sections))
	for _, s := range ds.Sections {
		sections = append(sections, sectionmapper.ModelSectionToProtoSection(s))
	}
	halls := make([]*proto.Hall, 0, len(ds.Halls))
	for _, h := range ds.Halls {
		hall, err := hallmapper.ModelHallToProtoHall(h)
		if err != nil {
			return nil, fmt.Errorf("unable to map hall: %w", err)
		}
		halls = append(halls, hall)
	}

	return &proto.DepartmentStore{
		Id:                            ds.ID,
		TradingPoint:                  tradingPoint,
		ReceiptingPointWithAccounting: receiptingPointWithAccounting,
		Sections:                      sections,
		Halls:                         halls,
	}, nil
}
