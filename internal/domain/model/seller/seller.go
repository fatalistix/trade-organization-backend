package seller

import (
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/place_of_work"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
)

type Seller struct {
	ID          int32
	FirstName   string
	LastName    string
	MiddleName  string
	BirthDate   core.Date
	Salary      core.Money
	PhoneNumber string
	WorksAt     *WorksAt
}

type WorksAt struct {
	PlaceOfWork  PlaceOfWork
	TradingPoint TradingPoint
}

type PlaceOfWork struct {
	ID   int32
	Type place_of_work.Type
}

type TradingPoint struct {
	ID   int32
	Type trading_point.Type
}
