package events

import "errors"

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func (p Priority) Validate() error {
	switch p {
	case PriorityHigh, PriorityMedium, PriorityLow:
		return nil
	default:
		return errors.New("invalid priority")
	}
}

func (p Priority) IsValid() bool {
	return p.Validate() == nil
}

func NewPriority(priority string) (Priority, error) {
	switch priority {
	case "low":
		return PriorityLow, nil
	case "medium":
		return PriorityMedium, nil
	case "high":
		return PriorityHigh, nil
	default:
		return "", errors.New("invalid priority")
	}
}
