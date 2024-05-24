package department_store

import (
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/hall"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/receipting_point_with_accounting"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/section"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
)

type DepartmentStore struct {
	ID                            int32
	TradingPoint                  trading_point.TradingPoint
	ReceiptingPointWithAccounting receipting_point_with_accounting.ReceiptingPointWithAccounting
	Sections                      []section.Section
	Halls                         []hall.Hall
}
