package place_of_work

import "fmt"

type Type string

const (
	TypeHall  Type = "hall"
	TypeKiosk Type = "kiosk"
	TypeTray  Type = "tray"
)

func TypeFromString(s string) (Type, error) {
	switch s {
	case string(TypeHall):
		return TypeHall, nil
	case string(TypeKiosk):
		return TypeKiosk, nil
	case string(TypeTray):
		return TypeTray, nil
	default:
		return "", fmt.Errorf("unknown type: %s", s)
	}
}
