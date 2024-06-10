package seller

import (
	"database/sql"
	"fmt"
	modelcore "github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/place_of_work"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/seller"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
	"github.com/fatalistix/trade-organization-backend/internal/repository/core"
	"github.com/fatalistix/trade-organization-backend/internal/repository/tradingpoint"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
)

type seller struct {
	ID               int32
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

	salary, err := modelcore.MoneyFromString(s.Salary)
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

	birthDate, err := modelcore.DateFromString(s.BirthDate)
	if err != nil {
		return model.Seller{}, fmt.Errorf("%s: unable to convert string to model date: %w", op, err)
	}

	return model.Seller{
		ID:          s.ID,
		FirstName:   s.FirstName,
		LastName:    s.LastName,
		MiddleName:  s.MiddleName,
		BirthDate:   birthDate,
		Salary:      salary,
		PhoneNumber: s.PhoneNumber,
		WorksAt:     worksAt,
	}, nil
}

func (s seller) ToProto() (*proto.Seller, error) {
	const op = "repository.seller.ToProto"

	birthDate, err := core.StringToProtoDate(s.BirthDate)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto date: %w", op, err)
	}

	salary, err := core.StringToProtoMoney(s.Salary)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert string to proto money: %w", op, err)
	}

	var worksAt *proto.WorksAt
	if s.PlaceOfWorkID.Valid && s.PlaceOfWorkType != nil && s.TradingPointID.Valid && s.TradingPointType != nil {
		tradingPointType, err := tradingpoint.StringToProtoTradingPointType(*s.TradingPointType)
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert string to proto type: %w", op, err)
		}
		placeOfWorkType, err := tradingpoint.StringToProtoPlaceOfWorkType(*s.PlaceOfWorkType)
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert string to proto type: %w", op, err)
		}

		worksAt = &proto.WorksAt{
			PlaceOfWork: &proto.PlaceOfWork{
				Id:   s.PlaceOfWorkID.Int32,
				Type: placeOfWorkType,
			},
			TradingPoint: &proto.TradingPoint{
				Id:   s.TradingPointID.Int32,
				Type: tradingPointType,
			},
		}
	}

	return &proto.Seller{
		Id:          s.ID,
		FirstName:   s.FirstName,
		LastName:    s.LastName,
		MiddleName:  s.MiddleName,
		BirthDate:   birthDate,
		Salary:      salary,
		PhoneNumber: s.PhoneNumber,
		WorksAt:     worksAt,
	}, nil
}
