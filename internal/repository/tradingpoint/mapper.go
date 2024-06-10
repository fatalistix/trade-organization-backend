package tradingpoint

import (
	"fmt"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
)

func ProtoTradingPointTypeToString(t proto.TradingPointType) (string, error) {
	switch t {
	case proto.TradingPointType_TRADING_POINT_STORE:
		return "store", nil
	case proto.TradingPointType_TRADING_POINT_DEPARTMENT_STORE:
		return "department_store", nil
	case proto.TradingPointType_TRADING_POINT_KIOSK:
		return "kiosk", nil
	case proto.TradingPointType_TRADING_POINT_TRAY:
		return "tray", nil
	default:
		return "", fmt.Errorf("unknown type: %s", t)
	}
}

func StringToProtoTradingPointType(t string) (proto.TradingPointType, error) {
	switch t {
	case "store":
		return proto.TradingPointType_TRADING_POINT_STORE, nil
	case "department_store":
		return proto.TradingPointType_TRADING_POINT_DEPARTMENT_STORE, nil
	case "kiosk":
		return proto.TradingPointType_TRADING_POINT_KIOSK, nil
	case "tray":
		return proto.TradingPointType_TRADING_POINT_TRAY, nil
	default:
		return proto.TradingPointType_TRADING_POINT_STORE, fmt.Errorf("unknown type: %s", t)
	}
}

func ProtoPlaceOfWorkTypeToString(t proto.PlaceOfWorkType) (string, error) {
	switch t {
	case proto.PlaceOfWorkType_PLACE_OF_WORK_HALL:
		return "hall", nil
	case proto.PlaceOfWorkType_PLACE_OF_WORK_KIOSK:
		return "kiosk", nil
	case proto.PlaceOfWorkType_PLACE_OF_WORK_TRAY:
		return "tray", nil
	default:
		return "", fmt.Errorf("unknown type: %s", t)
	}
}

func StringToProtoPlaceOfWorkType(t string) (proto.PlaceOfWorkType, error) {
	switch t {
	case "hall":
		return proto.PlaceOfWorkType_PLACE_OF_WORK_HALL, nil
	case "kiosk":
		return proto.PlaceOfWorkType_PLACE_OF_WORK_KIOSK, nil
	case "tray":
		return proto.PlaceOfWorkType_PLACE_OF_WORK_TRAY, nil
	default:
		return proto.PlaceOfWorkType_PLACE_OF_WORK_HALL, fmt.Errorf("unknown type: %s", t)
	}
}

func StringToProtoHallContainerType(t string) (proto.HallContainerType, error) {
	switch t {
	case "section":
		return proto.HallContainerType_HALL_CONTAINER_SECTION, nil
	case "store":
		return proto.HallContainerType_HALL_CONTAINER_STORE, nil
	default:
		return proto.HallContainerType_HALL_CONTAINER_SECTION, fmt.Errorf("unknown type: %s", t)
	}
}

func ProtoReceiptingPointWithAccountingTypeToString(t proto.ReceiptingPointWithAccountingType) (string, error) {
	switch t {
	case proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_DEPARTMENT_STORE:
		return "department_store", nil
	case proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_STORE:
		return "store", nil
	default:
		return "", fmt.Errorf("unknown type: %s", t)
	}
}

func StringToProtoReceiptingPointWithAccountingType(t string) (proto.ReceiptingPointWithAccountingType, error) {
	switch t {
	case "department_store":
		return proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_DEPARTMENT_STORE, nil
	case "store":
		return proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_STORE, nil
	default:
		return proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_DEPARTMENT_STORE, fmt.Errorf("unknown type: %s", t)
	}
}
