package result

import (
	"fmt"
	"open-feature/api/model"
	"open-feature/database"
)

type Rule struct {
	Type model.RuleType `json:"type"`
	Data any            `json:"data,omitempty"`
}

func (r *Rule) Result(v any) error {
	switch x := v.(type) {
	case database.StaticRule:
		r.Type = model.RuleTypeStatic
		r.Data = bool(x)
		return nil
	default:
		return fmt.Errorf("cannot map rule from type: %t", v)
	}
}
