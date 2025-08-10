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
