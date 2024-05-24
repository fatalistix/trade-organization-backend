package tradingpoint

import "fmt"

type Type string

const (
	TypeDepartmentStore Type = "department_store"
	TypeStore           Type = "store"
	TypeKiosk           Type = "kiosk"
	TypeTray            Type = "tray"
)

func TypeFromString(s string) (Type, error) {
	switch s {
	case string(TypeDepartmentStore):
		return TypeDepartmentStore, nil
	case string(TypeStore):
		return TypeStore, nil
	case string(TypeKiosk):
		return TypeKiosk, nil
	case string(TypeTray):
		return TypeTray, nil
	default:
		return "", fmt.Errorf("unknown type: %s", s)
	}
}
