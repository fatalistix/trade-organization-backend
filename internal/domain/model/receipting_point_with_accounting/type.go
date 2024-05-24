package receipting_point_with_accounting

import "fmt"

type Type string

const (
	TypeDepartmentStore Type = "department_store"
	TypeStore           Type = "store"
)

func (t Type) String() string {
	return string(t)
}

func TypeFromString(s string) (Type, error) {
	switch s {
	case string(TypeDepartmentStore):
		return TypeDepartmentStore, nil
	case string(TypeStore):
		return TypeStore, nil
	default:
		return "", fmt.Errorf("unknown type: %s", s)
	}
}
