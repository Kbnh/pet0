package domain

type Status int

const (
	StatusUnknown Status = iota
	StatusNew
	StatusInProgress
	StatusDone
)

func ParseStatus(s string) Status {
	switch s {
	case "new":
		return StatusNew
	case "in_progress":
		return StatusInProgress
	case "done":
		return StatusDone
	default:
		return StatusUnknown
	}
}
