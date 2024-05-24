package seller

import (
	"database/sql"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/place_of_work"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/seller"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
)

type seller struct {
	ID               int32
	Status           string
	FirstName        string
	LastName         string
	MiddleName       string
	BirthDate        string
	Salary           string
	PhoneNumber      string
	PlaceOfWorkID    sql.NullInt32
	PlaceOfWorkType  *string
	TradingPointID   sql.NullInt32
	TradingPointType *string
}

func (s seller) ToModel() (model.Seller, error) {
	const op = "repository.seller.ToModel"

	status, err := model.StatusFromString(s.Status)
	if err != nil {
		return model.Seller{}, fmt.Errorf("%s: unable to convert string to model status: %w", op, err)
	}

	salary, err := core.MoneyFromString(s.Salary)
	if err != nil {
		return model.Seller{}, fmt.Errorf("%s: unable to convert string to model money: %w", op, err)
	}

	var worksAt *model.WorksAt
	if s.PlaceOfWorkID.Valid && s.PlaceOfWorkType != nil && s.TradingPointID.Valid && s.TradingPointType != nil {
		placeOfWorkType, err := place_of_work.TypeFromString(*s.PlaceOfWorkType)
		if err != nil {
			return model.Seller{}, fmt.Errorf("%s: unable to convert string to model type: %w", op, err)
		}
		tradingPointType, err := trading_point.TypeFromString(*s.TradingPointType)
		if err != nil {
			return model.Seller{}, fmt.Errorf("%s: unable to convert string to model type: %w", op, err)
		}

		worksAt = &model.WorksAt{
			TradingPoint: model.TradingPoint{
				ID:   s.TradingPointID.Int32,
				Type: tradingPointType,
			},
			PlaceOfWork: model.PlaceOfWork{
				ID:   s.PlaceOfWorkID.Int32,
				Type: placeOfWorkType,
			},
		}
	} else {
		worksAt = nil
	}

	birthDate, err := core.DateFromString(s.BirthDate)
	if err != nil {
		return model.Seller{}, fmt.Errorf("%s: unable to convert string to model date: %w", op, err)
	}

	return model.Seller{
		ID:          s.ID,
		Status:      status,
		FirstName:   s.FirstName,
		LastName:    s.LastName,
		MiddleName:  s.MiddleName,
		BirthDate:   birthDate,
		Salary:      salary,
		PhoneNumber: s.PhoneNumber,
		WorksAt:     worksAt,
	}, nil
}
