package seller

import "fmt"

type Status string

const (
	StatusWorking    Status = "working"
	StatusNotWorking Status = "not_working"
)

func StatusFromString(s string) (Status, error) {
	switch s {
	case string(StatusWorking):
		return StatusWorking, nil
	case string(StatusNotWorking):
		return StatusNotWorking, nil
	default:
		return "", fmt.Errorf("unknown status: %s", s)
	}
}
