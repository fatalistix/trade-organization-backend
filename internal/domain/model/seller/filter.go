package seller

import (
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/place_of_work"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
)

type WorksAtFilterType string

const (
	WorksAtFilterTypeAll         WorksAtFilterType = "all"
	WorksAtFilterTypeOnlyNull    WorksAtFilterType = "only_null"
	WorksAtFilterTypeOnlyNotNull WorksAtFilterType = "only_not_null"
)

type Filter struct {
	WorksAtFilterType WorksAtFilterType
	TradingPointId    *int32
	TradingPointType  *trading_point.Type
	PlaceOfWorkId     *int32
	PlaceOfWorkType   *place_of_work.Type
	Search            *string
}
