package input

import (
	"open-feature/api/model"
)

type Rule struct {
	Type model.RuleType `json:"type"`
	Data any            `json:"data,omitempty"`
}
