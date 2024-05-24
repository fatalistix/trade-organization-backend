package section

import (
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/section"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
)

func ModelSectionToProtoSection(s section.Section) *proto.Section {
	return &proto.Section{
		Id:                s.ID,
		DepartmentStoreId: s.DepartmentStoreID,
	}
}
