package seller

import (
	"context"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	modelcore "github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/seller"
	"github.com/fatalistix/trade-organization-backend/internal/repository/core"
	"github.com/fatalistix/trade-organization-backend/internal/repository/tradingpoint"
	protocore "github.com/fatalistix/trade-organization-proto/gen/go/core"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
	"github.com/uptrace/bun"
	"log/slog"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(database *postgres.Database) *Repository {
	return &Repository{db: database.DB()}
}

func (r *Repository) CreateSellerContext(
	ctx context.Context,
	firstName string,
	lastName string,
	middleName string,
	birthDate modelcore.Date,
	salary modelcore.Money,
	phoneNumber string,
	worksAt *model.WorksAt,
) (int32, error) {
	const op = "repository.seller.RegisterSeller"

	var id int32
	values := map[string]interface{}{
		"first_name":   firstName,
		"last_name":    lastName,
		"middle_name":  middleName,
		"birth_date":   birthDate.String(),
		"salary":       salary.String(),
		"phone_number": phoneNumber,
	}
	if worksAt != nil {
		values["place_of_work_id"] = worksAt.PlaceOfWork.ID
		values["place_of_work_type"] = worksAt.PlaceOfWork.Type
		values["trading_point_id"] = worksAt.TradingPoint.ID
		values["trading_point_type"] = worksAt.TradingPoint.Type
	}

	err := r.db.NewInsert().
		Model(&values).
		TableExpr("seller").
		Returning("id").
		Scan(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new seller: %w", op, err)
	}

	return id, nil
}

func (r *Repository) SellersContext(ctx context.Context, filter *model.Filter) ([]model.Seller, error) {
	const op = "repository.seller.Sellers"

	query := r.db.NewSelect().
		Column("id", "first_name", "last_name", "middle_name", "birth_date", "salary", "phone_number",
			"place_of_work_id", "place_of_work_type", "trading_point_id", "trading_point_type",
		).TableExpr("seller")

	if filter.TradingPointId != nil {
		query = query.Where("trading_point_id = ?", filter.TradingPointId)
	}

	if filter.TradingPointType != nil {
		query = query.Where("trading_point_type = ?", filter.TradingPointType)
	}

	if filter.PlaceOfWorkId != nil {
		query = query.Where("place_of_work_id = ?", filter.PlaceOfWorkId)
	}

	if filter.PlaceOfWorkType != nil {
		query = query.Where("place_of_work_type = ?", filter.PlaceOfWorkType)
	}

	if filter.Search != nil {
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR middle_name ILIKE ?", "%"+*filter.Search+"%", "%"+*filter.Search+"%", "%"+*filter.Search+"%")
	}

	switch filter.WorksAtFilterType {
	case model.WorksAtFilterTypeOnlyNull:
		query = query.Where("place_of_work_id IS NULL AND trading_point_id IS NULL AND place_of_work_type IS NULL AND trading_point_type IS NULL")
	case model.WorksAtFilterTypeOnlyNotNull:
		query = query.Where("place_of_work_id IS NOT NULL AND trading_point_id IS NOT NULL AND place_of_work_type IS NOT NULL AND trading_point_type IS NOT NULL")
	case model.WorksAtFilterTypeAll:
		query = query.Where("place_of_work_id IS NOT NULL AND trading_point_id IS NOT NULL AND place_of_work_type IS NOT NULL AND trading_point_type IS NOT NULL")
		query = query.WhereOr("place_of_work_id IS NULL AND trading_point_id IS NULL AND place_of_work_type IS NULL AND trading_point_type IS NULL")
	}

	sellers := make([]*seller, 0)
	err := query.Scan(ctx, &sellers)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get all sellers: %w", op, err)
	}

	ms := make([]model.Seller, 0, len(sellers))
	for _, s := range sellers {
		m, err := s.ToModel()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert seller to model: %w", op, err)
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (r *Repository) SellerContext(ctx context.Context, id int32) (*proto.Seller, error) {
	const op = "repository.seller.Seller"

	var seller seller
	err := r.db.NewSelect().
		Column("id", "first_name", "last_name", "middle_name", "birth_date", "salary", "phone_number",
			"place_of_work_id", "place_of_work_type", "trading_point_id", "trading_point_type").
		Table("seller").
		Where("id = ?", id).
		Scan(ctx, &seller)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to select seller: %w", op, err)
	}

	return seller.ToProto()
}

func (r *Repository) PatchSellerContext(
	ctx context.Context,
	id int32,
	firstName *string,
	lastName *string,
	middleName *string,
	birthDate *protocore.Date,
	salary *protocore.Money,
	phoneNumber *string,
	worksAt *proto.NewWorksAt,
) error {
	const op = "repository.seller.UpdateSeller"

	values := map[string]interface{}{}
	if firstName != nil {
		values["first_name"] = *firstName
	}
	if lastName != nil {
		values["last_name"] = *lastName
	}
	if middleName != nil {
		values["middle_name"] = *middleName
	}
	if birthDate != nil {
		values["birth_date"] = core.ProtoDateToString(birthDate)
	}
	if salary != nil {
		values["salary"] = core.ProtoMoneyToString(salary)
	}
	if phoneNumber != nil {
		values["phone_number"] = *phoneNumber
	}
	if worksAt != nil {
		if worksAt.WorksAt == nil {
			values["place_of_work_id"] = nil
			values["place_of_work_type"] = nil
			values["trading_point_id"] = nil
			values["trading_point_type"] = nil
		} else {
			placeOfWorkType, err := tradingpoint.ProtoPlaceOfWorkTypeToString(worksAt.WorksAt.PlaceOfWork.Type)
			if err != nil {
				return fmt.Errorf("%s: unable to convert place of work type: %w", op, err)
			}

			tradingPointType, err := tradingpoint.ProtoTradingPointTypeToString(worksAt.WorksAt.TradingPoint.Type)
			if err != nil {
				return fmt.Errorf("%s: unable to convert trading point type: %w", op, err)
			}

			values["place_of_work_id"] = worksAt.WorksAt.PlaceOfWork.Id
			values["place_of_work_type"] = placeOfWorkType
			values["trading_point_id"] = worksAt.WorksAt.TradingPoint.Id
			values["trading_point_type"] = tradingPointType
		}
	}

	slog.Info(op, slog.Any("values", values))

	_, err := r.db.NewUpdate().
		Model(&values).
		TableExpr("seller").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to update seller: %w", op, err)
	}

	return nil
}
