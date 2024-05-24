package hall_container

import "fmt"

type Type string

const (
	TypeSection Type = "section"
	TypeStore   Type = "store"
)

func (t Type) String() string {
	return string(t)
}

func TypeFromString(s string) (Type, error) {
	switch s {
	case string(TypeSection):
		return TypeSection, nil
	case string(TypeStore):
		return TypeStore, nil
	default:
		return "", fmt.Errorf("unknown type: %s", s)
	}
}
