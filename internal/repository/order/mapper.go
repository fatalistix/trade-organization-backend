package order

import (
	"fmt"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
)

func StringToProtoOrderStatus(status string) (proto.OrderStatus, error) {
	switch status {
	case "in_progress":
		return proto.OrderStatus_ORDER_STATUS_IN_PROGRESS, nil
	case "completed":
		return proto.OrderStatus_ORDER_STATUS_COMPLETED, nil
	case "canceled":
		return proto.OrderStatus_ORDER_STATUS_CANCELLED, nil
	default:
		return proto.OrderStatus_ORDER_STATUS_IN_PROGRESS, fmt.Errorf("unknown status: %s", status)
	}
}

func ProtoOrderStatusToString(status proto.OrderStatus) (string, error) {
	switch status {
	case proto.OrderStatus_ORDER_STATUS_IN_PROGRESS:
		return "in_progress", nil
	case proto.OrderStatus_ORDER_STATUS_COMPLETED:
		return "completed", nil
	case proto.OrderStatus_ORDER_STATUS_CANCELLED:
		return "canceled", nil
	default:
		return "", fmt.Errorf("unknown status: %s", status)
	}
}
