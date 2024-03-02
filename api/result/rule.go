package result

import (
	"fmt"
	"open-feature/database"
)

type Rule struct {
	Type RuleType `json:"type"`
	Data any      `json:"data,omitempty"`
}

func (r *Rule) Map(v any) error {
	switch x := v.(type) {
	case database.StaticRule:
		r.Type = RuleTypeStatic
		r.Data = bool(x)
		return nil
	default:
		return fmt.Errorf("cannot map rule from type: %t", v)
	}
}

type RuleType string

const (
	RuleTypeStatic RuleType = "static"
)
