package supplier

import (
	"fmt"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/supplier"
)

func StringToProtoSupplierType(s string) (proto.SupplierType, error) {
	switch s {
	case "available":
		return proto.SupplierType_SUPPLIER_TYPE_AVAILABLE, nil
	case "not_available":
		return proto.SupplierType_SUPPLIER_TYPE_NOT_AVAILABLE, nil
	default:
		return 0, fmt.Errorf("unknown supplier type: %s", s)
	}
}

func ProtoSupplierTypeToString(t proto.SupplierType) (string, error) {
	switch t {
	case proto.SupplierType_SUPPLIER_TYPE_AVAILABLE:
		return "available", nil
	case proto.SupplierType_SUPPLIER_TYPE_NOT_AVAILABLE:
		return "not_available", nil
	default:
		return "", fmt.Errorf("unknown supplier type: %d", t)
	}
}
